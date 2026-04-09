import { useState, useEffect } from 'react';
import { differenceInDays, addDays, getTime, isSameDay, format } from 'date-fns';
import { DndContext, DragOverlay, pointerWithin, useSensor, useSensors, PointerSensor, KeyboardSensor } from '@dnd-kit/core';
import type { DragEndEvent, DragOverEvent, DragStartEvent } from '@dnd-kit/core';
import { arrayMove, sortableKeyboardCoordinates } from '@dnd-kit/sortable';
import toast from 'react-hot-toast';

import { CalendarDateItem } from './CalendarDateItem';
import { ScheduleWeekContainer } from '@/component/schedule/ScheduleWeekContainer';
import { ScheduleItem } from '@/component/schedule/ScheduleItem';

import { Type } from '@/type';
import { Model } from '@/model';
import { useDateKey } from '@/component/useDateKey';
import { useSchedule } from '@/provider/ScheduleProvider';
import { useSelection } from '@/provider/SelectionContext';

import { updateBulkSchedules } from '@/backend-api/updateBulkSchedules';
import { getColorClassName } from '@/component/calendar/color-module';
import { hasDateLabel } from '@/component/schedule/schedule-module';

type Props = {
  weeks: Date[][];
};

export const CalenderDates = ({ weeks }: Props) => {
  const { dateToKey } = useDateKey();
  const { masterSchedules, customSchedules, schedulesByType, setSchedulesByType, removeSchedule, saveSchedule, changeScheduleColor } =
    useSchedule();
  const { selectedIds, clearSelection, setAllSchedules } = useSelection();

  const [activeSchedule, setActiveSchedule] = useState<Type.Schedule | null>(null);
  const [bulkDragOrigins, setBulkDragOrigins] = useState<Map<string, Date>>(new Map());
  const [queueUpdateSortSchedules, setQueueUpdateSortSchedules] = useState<Type.Schedule[][]>([]);
  const sensors = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 5 } }),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  );

  // 範囲選択・同日チェック用にスケジュール一覧を日付順で管理する
  useEffect(() => {
    const allItems = [...masterSchedules, ...customSchedules];
    allItems.sort((a, b) => {
      if (a.date !== b.date) return a.date.localeCompare(b.date);
      return a.type === 'master' ? -1 : 1;
    });
    const refs = allItems.flatMap((item) =>
      [...item.schedules].sort((a, b) => a.order - b.order).map((s) => ({ id: s.id, date: item.date }))
    );
    setAllSchedules(refs);
  }, [masterSchedules, customSchedules, setAllSchedules]);

  const handleDragStart = (event: DragStartEvent) => {
    const { active } = event;
    const id = active.id.toString();
    const mSchedules = new Model.ScheduleDateItemList([...masterSchedules, ...customSchedules]);
    const schedule = mSchedules.getSchedule(id);

    if (!schedule) {
      return;
    }

    setActiveSchedule(schedule.toTypeSchedule());

    if (selectedIds.has(id) && selectedIds.size > 1) {
      // バルクドラッグ: 全選択アイテムの元 startDate をスナップショット
      const origins = new Map<string, Date>();
      for (const sid of selectedIds) {
        const s = mSchedules.getSchedule(sid);
        if (s) origins.set(sid, new Date(s.getStartDate()));
      }
      setBulkDragOrigins(origins);
    } else {
      // 未選択アイテムをドラッグ → 選択を解除して単体ドラッグ
      clearSelection();
      setBulkDragOrigins(new Map());
    }
  };

  const handleDragOver = (event: DragOverEvent) => {
    const { active, over } = event;

    if (!over) {
      return;
    }

    const beforeDate = active.data.current?.date;
    const afterDate = over.data.current?.date;
    const afterType: Type.ScheduleType | undefined = over.data.current?.type;

    if (!beforeDate || !afterDate || Number.isNaN(getTime(afterDate)) || !afterType) {
      return;
    }

    const mSchedules = new Model.ScheduleDateItemList([...masterSchedules, ...customSchedules]);
    const beforeSchedule = mSchedules.getSchedule(active.id.toString());

    if (!beforeSchedule) {
      return;
    }

    const beforeType: Type.ScheduleType = beforeSchedule.getType();

    if (beforeType === afterType) {
      const mSchedules = new Model.ScheduleDateItemList(schedulesByType(beforeType));

      const diff = differenceInDays(afterDate, beforeDate);
      const end = addDays(beforeSchedule.toTypeSchedule().endDate, diff);
      const afterSchedule: Type.Schedule = {
        ...beforeSchedule.toTypeSchedule(),
        startDate: afterDate,
        endDate: end,
      };
      mSchedules.removeSchedule(beforeSchedule.getId());

      const aDateKey = Model.ScheduleDateItem.toKey(afterDate);
      const mAfterType = new Model.ScheduleType(afterType);
      mSchedules.addSchedule(aDateKey, mAfterType, new Model.Schedule(afterSchedule));

      setSchedulesByType(beforeType, mSchedules.toTypeScheduleDateItems());

      const dateItem = mSchedules.getDateItem(aDateKey, mAfterType);
      if (dateItem) {
        const newSchedules = dateItem.toTypeSchedules();
        updateSortSchedules(newSchedules);
      }
      return;
    }

    const beforeSchedules = schedulesByType(beforeType);
    const afterSchedules = schedulesByType(afterType);
    const beforeMSchedules = new Model.ScheduleDateItemList(beforeSchedules);
    const afterMSchedules = new Model.ScheduleDateItemList(afterSchedules);

    beforeMSchedules.removeSchedule(beforeSchedule.getId());

    const aDateKey = Model.ScheduleDateItem.toKey(afterDate);
    const mAfterType = new Model.ScheduleType(afterType);
    const diff = differenceInDays(afterDate, beforeDate);
    const end = addDays(beforeSchedule.toTypeSchedule().endDate, diff);
    const afterSchedule: Type.Schedule = {
      id: beforeSchedule.getId(),
      name: beforeSchedule.getName(),
      startDate: afterDate,
      endDate: end,
      color: beforeSchedule.getColor(),
      type: afterType,
      order: 0,
    };
    afterMSchedules.addSchedule(aDateKey, mAfterType, new Model.Schedule(afterSchedule));

    setSchedulesByType(beforeType, beforeMSchedules.toTypeScheduleDateItems());
    setSchedulesByType(afterType, afterMSchedules.toTypeScheduleDateItems());

    const mBeforeType = new Model.ScheduleType(beforeType);
    const beforeDateItem = beforeMSchedules.getDateItem(beforeDate, mBeforeType);
    if (beforeDateItem) {
      const beforeNewSchedules = beforeDateItem.toTypeSchedules();
      updateSortSchedules(beforeNewSchedules);
    }

    const afterDateItem = afterMSchedules.getDateItem(aDateKey, mAfterType);
    if (afterDateItem) {
      const afterNewSchedules = afterDateItem.toTypeSchedules();
      updateSortSchedules(afterNewSchedules);
    }
  };

  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;

    if (!over) {
      setActiveSchedule(null);
      setBulkDragOrigins(new Map());
      return;
    }

    // バルクドラッグ
    if (bulkDragOrigins.size > 1) {
      handleBulkDragEnd(active.id.toString());
      return;
    }

    if (active.id === over.id) {
      setActiveSchedule(null);
      return;
    }

    const allSchedules = new Model.ScheduleDateItemList([...masterSchedules, ...customSchedules]);
    const beforeSchedule = allSchedules.getSchedule(active.id.toString());

    if (!beforeSchedule) {
      setActiveSchedule(null);
      return;
    }

    const dateKey = Model.ScheduleDateItem.toKey(beforeSchedule.getStartDate());
    const type = new Model.ScheduleType(beforeSchedule.getType());
    const mSchedules = new Model.ScheduleDateItemList(schedulesByType(type.String()));
    const dateItem = mSchedules.getDateItem(dateKey, type);

    if (!dateItem) {
      setActiveSchedule(null);
      return;
    }

    const beforeSchedules = dateItem.getSchedules();

    if (!beforeSchedules || beforeSchedules.length === 0) {
      setActiveSchedule(null);
      return;
    }

    const beforeIndex = beforeSchedules.findIndex((schedule) => schedule.getId() === active.id);
    const afterIndex = beforeSchedules.findIndex((schedule) => schedule.getId() === over.id);

    if (beforeIndex === -1 || afterIndex === -1) {
      setActiveSchedule(null);
      return;
    }

    const newSchedules = arrayMove(beforeSchedules, beforeIndex, afterIndex);
    mSchedules.setSchedules(dateKey, type, newSchedules);
    setSchedulesByType(type.String(), mSchedules.toTypeScheduleDateItems());
    setActiveSchedule(null);

    const newDateItem = mSchedules.getDateItem(dateKey, type);
    if (newDateItem) {
      const newSchedules = newDateItem.toTypeSchedules();
      updateSortSchedules(newSchedules);
    }
  };

  const handleBulkDragEnd = (primaryId: string) => {
    const allSchedules = new Model.ScheduleDateItemList([...masterSchedules, ...customSchedules]);
    const primaryCurrent = allSchedules.getSchedule(primaryId);
    const primaryOriginalStart = bulkDragOrigins.get(primaryId);

    if (!primaryCurrent || !primaryOriginalStart) {
      setActiveSchedule(null);
      setBulkDragOrigins(new Map());
      clearSelection();
      return;
    }

    const dateDelta = differenceInDays(primaryCurrent.getStartDate(), primaryOriginalStart);
    // DragOver 後のプライマリのタイプを移動先タイプとして全アイテムに適用する
    const targetType = primaryCurrent.getType();

    // 全選択アイテムの移動後スケジュールを構築（タイプは移動先に統一）
    const movedItems: Type.Schedule[] = [];

    for (const [id, originalStart] of bulkDragOrigins) {
      const s = allSchedules.getSchedule(id);
      if (!s) continue;

      // プライマリは DragOver で既に startDate/endDate が更新済みのため、
      // duration は s.getStartDate() を基準に計算する（originalStart だと二重シフトになる）
      const newStart = id === primaryId ? primaryCurrent.getStartDate() : addDays(originalStart, dateDelta);
      const duration = differenceInDays(s.getEndDate(), s.getStartDate());
      const newEnd = addDays(newStart, duration);

      movedItems.push({
        ...s.toTypeSchedule(),
        type: targetType,
        startDate: newStart,
        endDate: newEnd,
      });
    }

    // master・custom 両リストから選択アイテムを削除
    // （元タイプを問わず両方から削除することで、タイプ変更も自然に処理できる）
    const masterList = new Model.ScheduleDateItemList([...masterSchedules]);
    const customList = new Model.ScheduleDateItemList([...customSchedules]);
    for (const [id] of bulkDragOrigins) {
      masterList.removeSchedule(id);
      customList.removeSchedule(id);
    }

    // 移動先タイプのリストに追加
    const targetList = targetType === 'master' ? masterList : customList;
    const destDateKeys = new Set(movedItems.map((s) => Model.ScheduleDateItem.toKey(s.startDate)));
    for (const destDateKey of destDateKeys) {
      const mType = new Model.ScheduleType(targetType);
      const itemsForDate = movedItems
        .filter((s) => Model.ScheduleDateItem.toKey(s.startDate) === destDateKey)
        .sort((a, b) => a.order - b.order);
      for (const s of itemsForDate) {
        targetList.addSchedule(destDateKey, mType, new Model.Schedule(s));
      }

      // 移動先セルの全スケジュールを API へ
      const dateItem = targetList.getDateItem(destDateKey, mType);
      if (dateItem) {
        updateSortSchedules(dateItem.toTypeSchedules());
      }
    }

    setSchedulesByType('master', masterList.toTypeScheduleDateItems());
    setSchedulesByType('custom', customList.toTypeScheduleDateItems());

    setActiveSchedule(null);
    setBulkDragOrigins(new Map());
    clearSelection();
  };

  const updateSortSchedules = (targetSchedules: Type.Schedule[]) => {
    if (!targetSchedules || targetSchedules.length === 0) {
      return;
    }

    const newQueue = [...queueUpdateSortSchedules, targetSchedules];
    setQueueUpdateSortSchedules(newQueue);

    (async () => {
      const tss = newQueue.pop();
      if (!tss) {
        return;
      }

      try {
        await updateBulkSchedules(tss);
      } catch (e) {
        toast.error(String(e));
        setQueueUpdateSortSchedules(newQueue);
        return;
      }

      setQueueUpdateSortSchedules(newQueue);
    })();
  };

  const generateDragLabel = (schedule: Type.Schedule): string => {
    const dateFormat = 'M/d';
    if (!hasDateLabel(schedule)) {
      return schedule.name;
    }
    return `${schedule.name} (${format(schedule.startDate, dateFormat)} ~ ${format(schedule.endDate, dateFormat)})`;
  };

  return (
    <DndContext
      onDragStart={handleDragStart}
      onDragEnd={handleDragEnd}
      onDragOver={handleDragOver}
      collisionDetection={pointerWithin}
      sensors={sensors}
    >
      <div className="border-b border-r">
        {weeks.map((week, i) => {
          return (
            <div key={dateToKey(week[0]) + '-frame-week'} className="grid">
              <div data-date={dateToKey(week[0])} className=" calendar-item col-start-1 row-start-1 grid grid-cols-7">
                {week.map((date) => {
                  return (
                    <div key={dateToKey(date) + '-frame-date'} className="border-l border-t">
                      <CalendarDateItem date={date} />
                    </div>
                  );
                })}
              </div>
              <div className="col-start-1 row-start-1 mt-10 pb-1">
                <ScheduleWeekContainer week={week} activeSchedule={activeSchedule} />
              </div>
            </div>
          );
        })}
      </div>
      {activeSchedule && (
        <DragOverlay>
          <div className="relative">
            <div className={`flex touch-none items-center rounded px-1.5 py-1 ${getColorClassName(activeSchedule.color)}`}>
              <span className="line-clamp-2 text-[0.6rem] md:line-clamp-1 md:text-xs">{generateDragLabel(activeSchedule)}</span>
            </div>
            {bulkDragOrigins.size > 1 && (
              <span className="absolute -left-2 -top-2 flex min-w-max items-center justify-center rounded-full bg-blue-600 px-1.5 py-0.5 text-[0.6rem] font-bold text-white shadow">
                {bulkDragOrigins.size}件移動中
              </span>
            )}
          </div>
        </DragOverlay>
      )}
    </DndContext>
  );
};
