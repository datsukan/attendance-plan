import { parse, differenceInDays, addDays } from 'date-fns';
import { DndContext, pointerWithin, useDroppable, useSensor, useSensors, MouseSensor } from '@dnd-kit/core';
import type { DragEndEvent } from '@dnd-kit/core';

import { CalendarDateItem } from './CalendarDateItem';
import { ScheduleWeekContainer } from '@/component/schedule/ScheduleWeekContainer';

import { Schedule } from '@/type/schedule';
import { EditSchedule } from '@/model/edit-schedule';
import { dateKey } from './calendar-module';

type Props = {
  weeks: Date[][];
  schedules: Schedule[];
  removeSchedule: (id: string) => void;
  saveSchedule: (editSchedule: EditSchedule) => void;
  changeScheduleColor: (id: string, color: string) => void;
};

export const CalenderDates = ({ weeks, schedules, removeSchedule, saveSchedule, changeScheduleColor }: Props) => {
  const sensors = useSensors(useSensor(MouseSensor, { activationConstraint: { distance: 5 } }));

  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;

    if (!active || !over) {
      return;
    }

    const schedule = schedules.find((schedule) => schedule.id === (active.id as string));
    if (!schedule) {
      return;
    }

    const dateStr = over.data.current?.date;
    const type = over.data.current?.type;
    if (!dateStr || !type) {
      return;
    }

    const date = parse(dateStr, 'yyyy-MM-dd', new Date());
    const diff = differenceInDays(date, schedule.startDate);
    const start = date;
    const end = addDays(schedule.endDate, diff);
    const editSchedule = new EditSchedule(schedule);
    editSchedule.setStartDate(start);
    editSchedule.setEndDate(end);
    editSchedule.setType(type);

    saveSchedule(editSchedule);
  };

  return (
    <DndContext onDragEnd={handleDragEnd} collisionDetection={pointerWithin} sensors={sensors}>
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
