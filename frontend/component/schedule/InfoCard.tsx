import { format, isEqual } from 'date-fns';
import { ja } from 'date-fns/locale';

import { Schedule } from '@/type/schedule';
import { getColorClassName } from '@/component/calendar/color-module';
import { toScheduleTypeName } from '@/component/schedule/schedule-module';

type Props = {
  schedule: Schedule;
};

export const InfoCard = ({ schedule }: Props) => {
  return (
    <div className="rounded-lg shadow-lg bg-white overflow-hidden p-3 space-y-2">
      <div className="flex gap-2 items-center">
        <div className="px-2 py-1 rounded border text-sm">{toScheduleTypeName(schedule.type)}</div>
        <div className={`rounded-full size-5 ${getColorClassName(schedule.color)}`}></div>
        <div className="text-lg">{schedule.name}</div>
      </div>
      <div className="text-sm">
        {format(schedule.startDate, 'yyyy年M月d日(E)', { locale: ja })}
        {isEqual(schedule.startDate, schedule.endDate) ? '' : ` ~ ${format(schedule.endDate, 'yyyy年M月d日(E)', { locale: ja })}`}
      </div>
    </div>
  );
};
