import { format, isEqual } from 'date-fns';
import { ja } from 'date-fns/locale';

import { Type } from '@/type';
import { getColorClassName } from '@/component/calendar/color-module';
import { toScheduleTypeName } from '@/component/schedule/schedule-module';

type Props = {
  schedule: Type.Schedule;
};

export const InfoCard = ({ schedule }: Props) => {
  return (
    <div className="space-y-2 overflow-hidden rounded-lg bg-white p-3 shadow-lg">
      <div className="flex items-center gap-2">
        <div className="rounded border px-2 py-1 text-sm">{toScheduleTypeName(schedule.type)}</div>
        <div className={`size-5 rounded-full ${getColorClassName(schedule.color)}`}></div>
        <div className="text-lg">{schedule.name}</div>
      </div>
      <div className="text-sm">
        {format(schedule.startDate, 'yyyy年M月d日(E)', { locale: ja })}
        {isEqual(schedule.startDate, schedule.endDate) ? '' : ` ~ ${format(schedule.endDate, 'yyyy年M月d日(E)', { locale: ja })}`}
      </div>
    </div>
  );
};
