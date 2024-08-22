import { useEffect, useRef, useState } from 'react';

import { Schedule } from '@/type/schedule';
import { EditSchedule } from '@/model/edit-schedule';

import { ScheduleWeek } from '@/component/schedule/ScheduleWeek';

type Props = {
  index: number;
  week: Date[];
  masterSchedules: Schedule[];
  customSchedules: Schedule[];
  changeHeight: (index: number, height: number) => void;
  removeSchedule: (id: string) => void;
  saveSchedule: (editSchedule: EditSchedule) => void;
  changeScheduleColor: (id: string, color: string) => void;
};

export const ScheduleWeeks = ({
  index,
  week,
  changeHeight,
  masterSchedules,
  customSchedules,
  removeSchedule,
  saveSchedule,
  changeScheduleColor,
}: Props) => {
  const [show, setShow] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!ref.current) {
      return;
    }

    changeHeight(index, ref.current.clientHeight);
    setShow(true);
  }, [index, changeHeight]);

  return (
    <div ref={ref} className={`${show ? 'flex' : 'hidden'} flex-col gap-10`}>
      <ScheduleWeek
        dates={week}
        schedules={masterSchedules}
        removeSchedule={removeSchedule}
        saveSchedule={saveSchedule}
        changeScheduleColor={changeScheduleColor}
      />
      <ScheduleWeek
        dates={week}
        schedules={customSchedules}
        removeSchedule={removeSchedule}
        saveSchedule={saveSchedule}
        changeScheduleColor={changeScheduleColor}
      />
    </div>
  );
};
