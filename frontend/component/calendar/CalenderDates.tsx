import { useState } from 'react';
import { parse, differenceInDays, addDays, getTime } from 'date-fns';
import { DndContext, pointerWithin, useDroppable, useSensor, useSensors, MouseSensor, PointerSensor, KeyboardSensor } from '@dnd-kit/core';
import type { DragEndEvent, DragOverEvent, DragStartEvent } from '@dnd-kit/core';
import { arrayMove, SortableContext, sortableKeyboardCoordinates, useSortable } from '@dnd-kit/sortable';

import { CalendarDateItem } from './CalendarDateItem';
import { ScheduleWeekContainer } from '@/component/schedule/ScheduleWeekContainer';

import { Schedule } from '@/type/schedule';
import { EditSchedule } from '@/model/edit-schedule';
import { dateKey, toDate } from './calendar-module';

type Props = {
  weeks: Date[][];
  schedules: Schedule[];
  setSchedules: (schedules: Schedule[]) => void;
  removeSchedule: (id: string) => void;
  saveSchedule: (editSchedule: EditSchedule) => void;
  changeScheduleColor: (id: string, color: string) => void;
};

export const CalenderDates = ({ weeks, schedules, setSchedules, removeSchedule, saveSchedule, changeScheduleColor }: Props) => {
  const [activeSchedule, setActiveSchedule] = useState<Schedule | null>(null);
  const sensors = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 0.1 } }),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  );

  const handleDragStart = (event: DragStartEvent) => {
    const { active } = event;
    const id = active.id;
    const schedule = schedules.find((schedule) => schedule.id === id);

    if (!schedule) {
      return;
    }

    setActiveSchedule(schedule);
  };

  const handleDragOver = (event: DragOverEvent) => {
    const { active, over } = event;

    if (!over) {
      return;
    }

    const beforeDate = active.data.current?.date;
    const afterDate = over.data.current?.date;
    const afterType = over.data.current?.type;

    if (!beforeDate || !afterDate || Number.isNaN(getTime(afterDate)) || !afterType) {
      return;
    }

    const newSchedules = schedules.map((schedule) => {
      if (schedule.id === active.id) {
        const diff = differenceInDays(afterDate, beforeDate);
        const end = addDays(schedule.endDate, diff);
        return { ...schedule, startDate: afterDate, endDate: end, type: afterType };
      }

      return schedule;
    });

    setSchedules(newSchedules);
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

    const beforeIndex = schedules.findIndex((schedule) => schedule.id === active.id);
    const afterIndex = schedules.findIndex((schedule) => schedule.id === over.id);
    if (beforeIndex === -1 || afterIndex === -1) {
      setActiveSchedule(null);
      return;
    }

    const newSchedules = arrayMove(schedules, beforeIndex, afterIndex);
    setSchedules(newSchedules);
    setActiveSchedule(null);
  };

  return (
    <DndContext
      onDragStart={handleDragStart}
      onDragEnd={handleDragEnd}
      onDragOver={handleDragOver}
      collisionDetection={pointerWithin}
      sensors={sensors}
    >
      <div className="border-r border-b">
        {weeks.map((week, i) => {
          return (
            <div key={dateKey(week[0]) + '-frame-week'} className="grid">
              <div data-date={dateKey(week[0])} className=" col-start-1 row-start-1 grid grid-cols-7 calendar-item">
                {week.map((date) => {
                  return (
                    <div key={dateKey(date) + '-frame-date'} className="border-t border-l">
                      <CalendarDateItem date={date} />
                    </div>
                  );
                })}
              </div>
              <div className="mt-10 pb-2 col-start-1 row-start-1">
                <ScheduleWeekContainer
                  week={week}
                  masterSchedules={schedules.filter((schedule) => schedule.type === 'master')}
                  customSchedules={schedules.filter((schedule) => schedule.type === 'custom')}
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
