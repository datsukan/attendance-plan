import { YearMonthLabel } from './YearMonthLabel';
import { AddScheduleButton } from './AddScheduleButton';
import { MoveMonthNav } from './MoveMonthNav';
import { DayOfWeeks } from './DayOfWeeks';

import { CreateSchedule } from '@/model/createSchedule';

type Props = {
  year: number;
  month: number;
  prev: () => void;
  next: () => void;
  reset: () => void;
  create: (createSchedule: CreateSchedule) => Promise<void>;
};

export const Header = ({ year, month, prev, next, reset, create }: Props) => {
  return (
    <>
      <div className="flex items-center justify-between gap-6 py-6">
        <div className="flex items-center gap-6">
          <YearMonthLabel year={year} month={month} />
          <AddScheduleButton create={create} />
        </div>
        <MoveMonthNav prev={prev} next={next} reset={reset} />
      </div>
      <DayOfWeeks />
    </>
  );
};
