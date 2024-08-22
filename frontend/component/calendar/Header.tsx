import { YearMonthLabel } from './YearMonthLabel';
import { AddMasterScheduleButton } from './AddMasterScheduleButton';
import { MoveMonthNav } from './MoveMonthNav';
import { DayOfWeeks } from './DayOfWeeks';

import { CreateSchedule } from '@/model/create-schedule';

type Props = {
  year: number;
  month: number;
  prev: () => void;
  next: () => void;
  reset: () => void;
  create: (createSchedule: CreateSchedule) => void;
};

export const Header = ({ year, month, prev, next, reset, create }: Props) => {
  return (
    <>
      <div className="py-6 flex gap-6 items-center justify-between">
        <div className="flex gap-6 items-center">
          <YearMonthLabel year={year} month={month} />
          <AddMasterScheduleButton create={create} />
        </div>
        <MoveMonthNav prev={prev} next={next} reset={reset} />
      </div>
      <DayOfWeeks />
    </>
  );
};
