import { Schedule } from '@/type/schedule';
import { EditSchedule } from '@/model/edit-schedule';

import { ScheduleWeek } from '@/component/schedule/ScheduleWeek';

type Props = {
  week: Date[];
  masterSchedules: Schedule[];
  customSchedules: Schedule[];
  removeSchedule: (id: string) => void;
  saveSchedule: (editSchedule: EditSchedule) => void;
  changeScheduleColor: (id: string, color: string) => void;
};

export const ScheduleWeekContainer = ({
  week,
  masterSchedules,
  customSchedules,
  removeSchedule,
  saveSchedule,
  changeScheduleColor,
}: Props) => {
  return (
    <div className="flex gap-6 flex-col min-h-full">
      <ScheduleWeek
        type="master"
        dates={week}
        schedules={masterSchedules}
        removeSchedule={removeSchedule}
        saveSchedule={saveSchedule}
        changeScheduleColor={changeScheduleColor}
      />
      <ScheduleWeek
        type="custom"
        dates={week}
        schedules={customSchedules}
        removeSchedule={removeSchedule}
        saveSchedule={saveSchedule}
        changeScheduleColor={changeScheduleColor}
      />
    </div>
  );
};
