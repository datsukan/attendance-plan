import { format } from 'date-fns';

import type { Schedule } from '@/type/schedule';

import { hasDateLabel } from '@/component/schedule/schedule-module';

type Props = {
  schedule: Schedule;
};

export const ScheduleItem = ({ schedule }: Props) => {
  const dateFormat = 'M/d';

  function generateLabel(): string {
    if (!hasDateLabel(schedule)) {
      return schedule.name;
    }

    return `${schedule.name} (${format(schedule.startDate, dateFormat)} ~ ${format(schedule.endDate, dateFormat)})`;
  }

  return (
    <div className={`px-2 h-6 flex items-center rounded ${schedule.styleClassName}`}>
      <span className="text-xs truncate">{generateLabel()}</span>
    </div>
  );
};
