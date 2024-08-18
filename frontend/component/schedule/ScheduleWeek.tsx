import { Schedule } from '@/type/schedule';

import { dateKey } from '@/component/calendar/calendar-module';
import { isShowItem, isDisplaySchedule, getColStartClassName, getColEndClassName } from '@/component/schedule/schedule-module';

import { ScheduleItem } from '@/component/schedule/ScheduleItem';

type Props = {
  dates: Date[];
  schedules: Schedule[];
};

export const ScheduleWeek = ({ dates, schedules }: Props) => {
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
    <div className="flex flex-col gap-1">
      <div className="grid grid-cols-7 gap-y-1">
        {displaySchedules.map((schedule) => {
          return dates.map((date, index) => {
            if (!isShowItem(index, schedule, date)) {
              return null;
            }

            const colStartClassName = getColStartClassName(index);
            const colEndClassName = getColEndClassName(index, schedule, dates);

            return (
              <div key={dateKey(date) + '-schedules'} className={`${colStartClassName} ${colEndClassName} pr-2`}>
                <ScheduleItem key={`${schedule.id}-${dateKey(date)}`} schedule={schedule} />
              </div>
            );
          });
        })}
      </div>
    </div>
  );
};
