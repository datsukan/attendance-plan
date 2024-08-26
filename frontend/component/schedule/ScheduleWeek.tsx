import { useDroppable } from '@dnd-kit/core';

import { ScheduleWeekItem } from '@/component/schedule/ScheduleWeekItem';

import { Schedule } from '@/type/schedule';
import { dateKey } from '@/component/calendar/calendar-module';
import { isShowItem, isDisplaySchedule, getColStartClassName, getColEndClassName } from '@/component/schedule/schedule-module';
import { EditSchedule } from '@/model/edit-schedule';

type Props = {
  type: 'master' | 'custom';
  dates: Date[];
  schedules: Schedule[];
  removeSchedule: (id: string) => void;
  saveSchedule: (editSchedule: EditSchedule) => void;
  changeScheduleColor: (id: string, color: string) => void;
};

export const ScheduleWeek = ({ type, dates, schedules, removeSchedule, saveSchedule, changeScheduleColor }: Props) => {
  if (!dates || dates.length === 0) {
    return;
  }

  if (!schedules || schedules.length === 0) {
    return;
  }

  const displaySchedules = schedules.filter((schedule) => {
    for (const date of dates) {
      if (isDisplaySchedule(schedule, date)) {
        return true;
      }
    }

    return false;
  });

  return (
    <div className="h-full min-h-10 grid">
      <div className="col-start-1 row-start-1 grid grid-cols-7 h-full">
        {dates.map((date) => (
          <Droppable key={`${type}-${dateKey(date)}`} id={`${type}-${dateKey(date)}`} date={dateKey(date)} type={type} />
        ))}
      </div>
      <div className="col-start-1 row-start-1 grid grid-cols-7 gap-y-1 grid-flow-col">
        {displaySchedules.map((schedule) => {
          let countSchedule = 0;
          return dates.map((date, index) => {
            if (!isShowItem(index, schedule, date)) {
              return null;
            }
            countSchedule++;

            const colStartClassName = getColStartClassName(index);
            const colEndClassName = getColEndClassName(index, schedule, dates);

            return (
              <div key={schedule.id} className={`pr-2 ${colStartClassName} ${colEndClassName}`}>
                <ScheduleWeekItem
                  schedule={schedule}
                  removeSchedule={removeSchedule}
                  saveSchedule={saveSchedule}
                  changeScheduleColor={changeScheduleColor}
                />
              </div>
            );
          });
        })}
      </div>
    </div>
  );
};

type DroppableProps = {
  id: string;
  date: string;
  type: 'master' | 'custom';
  children?: React.ReactNode;
};

const Droppable = ({ id, date, type, children }: DroppableProps) => {
  const { isOver, setNodeRef } = useDroppable({
    id: id,
    data: { date: date, type: type },
  });

  return (
    <div ref={setNodeRef} className={isOver ? 'bg-blue-50' : ''}>
      {children}
    </div>
  );
};
