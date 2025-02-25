import { Type } from '@/type';
import { ScheduleWeek } from '@/component/schedule/ScheduleWeek';
import { ScheduleTypeMaster, ScheduleTypeCustom } from '@/const/schedule';
import { useSchedule } from '@/provider/ScheduleProvider';

type Props = {
  week: Date[];
  activeSchedule: Type.Schedule | null;
};

export const ScheduleWeekContainer = ({ week, activeSchedule }: Props) => {
  const { masterSchedules, customSchedules } = useSchedule();

  return (
    <div className="flex min-h-full flex-col">
      <ScheduleWeek type={ScheduleTypeMaster} dates={week} schedules={masterSchedules} activeSchedule={activeSchedule} />
      <ScheduleWeek type={ScheduleTypeCustom} dates={week} schedules={customSchedules} activeSchedule={activeSchedule} />
    </div>
  );
};
