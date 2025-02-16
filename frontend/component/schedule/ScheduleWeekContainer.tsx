import { Type } from '@/type';
import { Model } from '@/model';

import { ScheduleWeek } from '@/component/schedule/ScheduleWeek';
import { ScheduleTypeMaster, ScheduleTypeCustom } from '@/const/schedule';

type Props = {
  week: Date[];
  masterSchedules: Type.ScheduleDateItem[];
  customSchedules: Type.ScheduleDateItem[];
  activeSchedule: Type.Schedule | null;
  removeSchedule: (id: string, type: Type.ScheduleType) => void;
  saveSchedule: (editSchedule: Model.EditSchedule) => void;
  changeScheduleColor: (id: string, type: Type.ScheduleType, color: string) => void;
};

export const ScheduleWeekContainer = ({
  week,
  masterSchedules,
  customSchedules,
  activeSchedule,
  removeSchedule,
  saveSchedule,
  changeScheduleColor,
}: Props) => {
  return (
    <div className="flex min-h-full flex-col">
      <ScheduleWeek
        type={ScheduleTypeMaster}
        dates={week}
        schedules={masterSchedules}
        activeSchedule={activeSchedule}
        removeSchedule={removeSchedule}
        saveSchedule={saveSchedule}
        changeScheduleColor={changeScheduleColor}
      />
      <ScheduleWeek
        type={ScheduleTypeCustom}
        dates={week}
        schedules={customSchedules}
        activeSchedule={activeSchedule}
        removeSchedule={removeSchedule}
        saveSchedule={saveSchedule}
        changeScheduleColor={changeScheduleColor}
      />
    </div>
  );
};
