import { useEffect, useRef, useState } from 'react';

import { Schedule } from '@/type/schedule';

import { ScheduleWeek } from '@/component/schedule/ScheduleWeek';

import { getMasterSchedules, getCustomSchedules } from '@/component/schedule/schedule-module';

type Props = {
  index: number;
  week: Date[];
  masterSchedules: Schedule[];
  customSchedules: Schedule[];
  changeHeight: (index: number, height: number) => void;
};

export const ScheduleWeeks = ({ index, week, changeHeight, masterSchedules, customSchedules }: Props) => {
  const [show, setShow] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (ref.current) {
      changeHeight(index, ref.current.clientHeight);
      setShow(true);
    }
  }, [index, changeHeight]);

  return (
    <div ref={ref} className={`${show ? 'flex' : 'hidden'} flex-col gap-10`}>
      <ScheduleWeek dates={week} schedules={masterSchedules} />
      <ScheduleWeek dates={week} schedules={customSchedules} />
    </div>
  );
};
