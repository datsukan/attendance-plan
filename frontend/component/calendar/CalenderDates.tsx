import { useState } from 'react';
import { differenceInDays, addDays, getTime, isSameDay } from 'date-fns';
import { DndContext, pointerWithin, useSensor, useSensors, PointerSensor, KeyboardSensor } from '@dnd-kit/core';
import type { DragEndEvent, DragOverEvent, DragStartEvent } from '@dnd-kit/core';
import { arrayMove, sortableKeyboardCoordinates } from '@dnd-kit/sortable';

import { CalendarDateItem } from './CalendarDateItem';
import { ScheduleWeekContainer } from '@/component/schedule/ScheduleWeekContainer';

import { Type } from '@/type';
import { Model } from '@/model';
import { useDateKey } from '@/component/useDateKey';

import { updateBulkSchedules } from '@/api/updateBulkSchedules';

type Props = {
  weeks: Date[][];
  masterSchedules: Type.ScheduleDateItem[];
  customSchedules: Type.ScheduleDateItem[];
  setMasterSchedules: (schedules: Type.ScheduleDateItem[]) => void;
  setCustomSchedules: (schedules: Type.ScheduleDateItem[]) => void;
  schedulesByType: (type: Type.ScheduleType) => Type.ScheduleDateItem[];
  setSchedulesByType: (type: Type.ScheduleType, schedules: Type.ScheduleDateItem[]) => void;
  removeSchedule: (id: string, type: Type.ScheduleType) => void;
  saveSchedule: (editSchedule: Model.EditSchedule) => void;
  changeScheduleColor: (id: string, type: Type.ScheduleType, color: string) => void;
};

export const CalenderDates = ({
  weeks,
  masterSchedules,
  customSchedules,
  setMasterSchedules,
  setCustomSchedules,
  schedulesByType,
  setSchedulesByType,
  removeSchedule,
  saveSchedule,
  changeScheduleColor,
}: Props) => {
  const { dateToKey } = useDateKey();

  const [activeSchedule, setActiveSchedule] = useState<Type.Schedule | null>(null);
  const [isRunningUpdateSortSchedules, setIsRunningUpdateSortSchedules] = useState(false);
  const sensors = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 0.1 } }),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  );

  const handleDragStart = (event: DragStartEvent) => {
    const { active } = event;
    const id = active.id;
    const mSchedules = new Model.ScheduleDateItemList([...masterSchedules, ...customSchedules]);
    const schedule = mSchedules.getSchedule(id.toString());

    if (!schedule) {
      return;
    }

    setActiveSchedule(schedule.toTypeSchedule());
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

  const updateSortSchedules = (targetSchedules: Type.Schedule[]) => {
    if (isRunningUpdateSortSchedules) {
      setTimeout(() => updateSortSchedules(targetSchedules), 100);
    }

    setIsRunningUpdateSortSchedules(true);

    if (!targetSchedules || targetSchedules.length === 0) {
      return;
    }

    (async () => {
      updateBulkSchedules(targetSchedules);
      setIsRunningUpdateSortSchedules(false);
    })();
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
                <ScheduleWeekContainer
                  week={week}
                  masterSchedules={masterSchedules}
                  customSchedules={customSchedules}
                  activeSchedule={activeSchedule}
                  removeSchedule={removeSchedule}
                  saveSchedule={saveSchedule}
                  changeScheduleColor={changeScheduleColor}
                />
              </div>
            </div>
          );
        })}
      </div>
    </DndContext>
  );
};
