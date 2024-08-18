import { YearMonthLabel } from './YearMonthLabel';
import { MoveMonthNav } from './MoveMonthNav';
import { DayOfWeeks } from './DayOfWeeks';

type Props = {
  year: number;
  month: number;
  prev: () => void;
  next: () => void;
  reset: () => void;
};

export const Header = ({ year, month, prev, next, reset }: Props) => {
  return (
    <>
      <div className="py-6 flex gap-6 items-center justify-between">
        <YearMonthLabel year={year} month={month + 1} />
        <MoveMonthNav prev={prev} next={next} reset={reset} />
      </div>
      <DayOfWeeks />
    </>
  );
};
