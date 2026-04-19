import { useState, useEffect } from 'react';
import { differenceInDays, addDays, getTime, isSameDay, format } from 'date-fns';
import { DndContext, DragOverlay, pointerWithin, useSensor, useSensors, PointerSensor } from '@dnd-kit/core';
import type { DragEndEvent, DragOverEvent, DragStartEvent } from '@dnd-kit/core';
import { arrayMove } from '@dnd-kit/sortable';
import toast from 'react-hot-toast';

import { CalendarDateItem } from './CalendarDateItem';
import { ScheduleWeekContainer } from '@/component/schedule/ScheduleWeekContainer';
import { ScheduleItem } from '@/component/schedule/ScheduleItem';

import { Type } from '@/type';
import { Model } from '@/model';
import { useDateKey } from '@/component/useDateKey';
import { useSchedule } from '@/provider/ScheduleProvider';
import { useSelection } from '@/provider/SelectionContext';
import { useUndo } from '@/provider/UndoProvider';

import { updateBulkSchedules } from '@/backend-api/updateBulkSchedules';
import { SessionExpiredError } from '@/backend-api/error';
import { getColorClassName } from '@/component/calendar/color-module';
import { hasDateLabel } from '@/component/schedule/schedule-module';

type Props = {
  weeks: Date[][];
};

export const CalenderDates = ({ weeks }: Props) => {
  const { dateToKey } = useDateKey();
  const { masterSchedules, customSchedules, schedulesByType, setSchedulesByType, setSchedulesByTypeFunctional, removeSchedule, saveSchedule, changeScheduleColor } =
    useSchedule();
  const { selectedIds, clearSelection, setAllSchedules } = useSelection();
  const { setUndoCommand } = useUndo();

  const [activeSchedule, setActiveSchedule] = useState<Type.Schedule | null>(null);
  const [bulkDragOrigins, setBulkDragOrigins] = useState<Map<string, Date>>(new Map());
  const [dragSnapshot, setDragSnapshot] = useState<Type.Schedule[]>([]);
  const [queueUpdateSortSchedules, setQueueUpdateSortSchedules] = useState<Type.Schedule[][]>([]);
  const sensors = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 5 } })
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
      const snapshotSchedules: Type.Schedule[] = [];
      for (const sid of selectedIds) {
        const s = mSchedules.getSchedule(sid);
        if (s) {
          origins.set(sid, new Date(s.getStartDate()));
          snapshotSchedules.push(s.toTypeSchedule());
        }
      }
      setBulkDragOrigins(origins);
      setDragSnapshot(snapshotSchedules);
    } else {
      // 未選択アイテムをドラッグ → 選択を解除して単体ドラッグ
      clearSelection();
      setBulkDragOrigins(new Map());

      // 単体ドラッグ: 移動元セルのスケジュール全体をスナップショット（並び替えにも対応）
      const sourceDateKey = Model.ScheduleDateItem.toKey(schedule.getStartDate());
      const sourceType = new Model.ScheduleType(schedule.getType());
      const sourceDateItem = mSchedules.getDateItem(sourceDateKey, sourceType);
      const cellSchedules = sourceDateItem ? sourceDateItem.toTypeSchedules() : [schedule.toTypeSchedule()];
      setDragSnapshot(cellSchedules);
    }
  };

  const handleDragOver = (event: DragOverEvent) => {
    const { active, over } = event;

    if (!over) {
      return;
    }

    const afterDate = over.data.current?.date;
    const afterType: Type.ScheduleType | undefined = over.data.current?.type;

    if (!afterDate || Number.isNaN(getTime(afterDate)) || !afterType) {
      return;
    }

    const mSchedules = new Model.ScheduleDateItemList([...masterSchedules, ...customSchedules]);
    const beforeSchedule = mSchedules.getSchedule(active.id.toString());

    if (!beforeSchedule) {
      return;
    }

    // active.data.current?.date は SortableContext を跨いで移動する際に
    // アイテムの unmount/remount によって古い値のまま残ることがある。
    // state から取得した startDate を使うことで常に最新の位置を参照する。
    const beforeDate = beforeSchedule.getStartDate();
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
    const beforeDateItem = beforeMSchedules.getDateItem(Model.ScheduleDateItem.toKey(beforeDate), mBeforeType);
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
      setDragSnapshot([]);
      return;
    }

    // バルクドラッグ
    if (bulkDragOrigins.size > 1) {
      handleBulkDragEnd(active.id.toString());
      return;
    }

    if (active.id === over.id) {
      // handleDragOver により既に別の日付/タイプに移動済みの可能性があるため変更有無を確認する
      if (dragSnapshot.length > 0 && activeSchedule) {
        const snapshotItem = dragSnapshot.find((s) => s.id === active.id.toString());
        const allSchedules = new Model.ScheduleDateItemList([...masterSchedules, ...customSchedules]);
        const currentSchedule = allSchedules.getSchedule(active.id.toString());
        if (
          snapshotItem &&
          currentSchedule &&
          (!isSameDay(currentSchedule.getStartDate(), snapshotItem.startDate) ||
            currentSchedule.getType() !== snapshotItem.type)
        ) {
          registerDragUndo(`「${activeSchedule.name}」を移動しました`, dragSnapshot);
        }
      }
      setActiveSchedule(null);
      setDragSnapshot([]);
      return;
    }

    const allSchedules = new Model.ScheduleDateItemList([...masterSchedules, ...customSchedules]);
    const beforeSchedule = allSchedules.getSchedule(active.id.toString());

    if (!beforeSchedule) {
      setActiveSchedule(null);
      setDragSnapshot([]);
      return;
    }

    const dateKey = Model.ScheduleDateItem.toKey(beforeSchedule.getStartDate());
    const type = new Model.ScheduleType(beforeSchedule.getType());
    const mSchedules = new Model.ScheduleDateItemList(schedulesByType(type.String()));
    const dateItem = mSchedules.getDateItem(dateKey, type);

    if (!dateItem) {
      setActiveSchedule(null);
      setDragSnapshot([]);
      return;
    }

    const beforeSchedules = dateItem.getSchedules();

    if (!beforeSchedules || beforeSchedules.length === 0) {
      setActiveSchedule(null);
      setDragSnapshot([]);
      return;
    }

    const beforeIndex = beforeSchedules.findIndex((schedule) => schedule.getId() === active.id);
    const afterIndex = beforeSchedules.findIndex((schedule) => schedule.getId() === over.id);

    if (beforeIndex === -1 || afterIndex === -1) {
      // 日付/タイプ変更ドラッグ（handleDragOver で処理済み）
      // 元の位置に戻った場合はトーストを出さない
      if (dragSnapshot.length > 0 && activeSchedule) {
        const snapshotItem = dragSnapshot.find((s) => s.id === active.id.toString());
        const currentSchedule = allSchedules.getSchedule(active.id.toString());
        if (
          snapshotItem &&
          currentSchedule &&
          (!isSameDay(currentSchedule.getStartDate(), snapshotItem.startDate) ||
            currentSchedule.getType() !== snapshotItem.type)
        ) {
          registerDragUndo(`「${activeSchedule.name}」を移動しました`, dragSnapshot);
        }
      }
      setActiveSchedule(null);
      setDragSnapshot([]);
      return;
    }

    const newSchedules = arrayMove(beforeSchedules, beforeIndex, afterIndex);
    newSchedules.forEach((schedule, i) => schedule.setOrder(i + 1));
    mSchedules.setSchedules(dateKey, type, newSchedules);
    setSchedulesByType(type.String(), mSchedules.toTypeScheduleDateItems());

    // 同じ位置に戻った場合はトーストを出さない
    if (dragSnapshot.length > 0 && activeSchedule && beforeIndex !== afterIndex) {
      registerDragUndo(`「${activeSchedule.name}」を移動しました`, dragSnapshot);
    }
    setActiveSchedule(null);
    setDragSnapshot([]);

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
      setDragSnapshot([]);
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

    // 元の位置に戻った場合はトーストを出さない
    const primarySnapshot = dragSnapshot.find((s) => s.id === primaryId);
    const hasActuallyMoved = dateDelta !== 0 || (primarySnapshot && targetType !== primarySnapshot.type);
    if (dragSnapshot.length > 0 && hasActuallyMoved) {
      registerDragUndo(`${dragSnapshot.length}件のスケジュールを移動しました`, dragSnapshot);
    }

    setActiveSchedule(null);
    setBulkDragOrigins(new Map());
    setDragSnapshot([]);
    clearSelection();
  };

  const registerDragUndo = (label: string, snapshot: Type.Schedule[]) => {
    const sorted = [...snapshot].sort((a, b) => a.order - b.order);
    const snapshotIds = new Set(snapshot.map((s) => s.id));
    const masterOriginals = sorted.filter((s) => s.type === 'master');
    const customOriginals = sorted.filter((s) => s.type !== 'master');

    setUndoCommand({
      label,
      execute: async () => {
        await updateBulkSchedules(sorted);

        setSchedulesByTypeFunctional('master', (prev) => {
          const list = new Model.ScheduleDateItemList(prev);
          snapshotIds.forEach((id) => list.removeSchedule(id));
          masterOriginals.forEach((s) => {
            list.addSchedule(
              Model.ScheduleDateItem.toKey(s.startDate),
              new Model.ScheduleType(s.type),
              new Model.Schedule(s)
            );
          });
          return list.toTypeScheduleDateItems();
        });

        setSchedulesByTypeFunctional('custom', (prev) => {
          const list = new Model.ScheduleDateItemList(prev);
          snapshotIds.forEach((id) => list.removeSchedule(id));
          customOriginals.forEach((s) => {
            list.addSchedule(
              Model.ScheduleDateItem.toKey(s.startDate),
              new Model.ScheduleType(s.type),
              new Model.Schedule(s)
            );
          });
          return list.toTypeScheduleDateItems();
        });
      },
    });
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
        if (e instanceof SessionExpiredError) return;
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
