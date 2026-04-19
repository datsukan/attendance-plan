import Link from 'next/link';
import { QuestionMarkCircleIcon } from '@heroicons/react/24/outline';

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
  create: (createSchedule: CreateSchedule[]) => Promise<void>;
};

export const Header = ({ year, month, prev, next, reset, create }: Props) => {
  return (
    <>
      <div className="flex flex-col items-center justify-between gap-6 py-6 md:flex-row">
        <div className="flex items-center gap-6">
          <YearMonthLabel year={year} month={month} />
          <AddScheduleButton create={create} />
        </div>
        <div className="flex items-center gap-3">
          <Link
            href="/guide"
            className="flex items-center gap-1 rounded-md px-3 py-1.5 text-sm text-gray-500 hover:bg-gray-100 hover:text-gray-700"
          >
            <QuestionMarkCircleIcon className="size-4" />
            <span className="mb-0.5">機能・使い方</span>
          </Link>
          <MoveMonthNav prev={prev} next={next} reset={reset} />
        </div>
      </div>
      <DayOfWeeks />
    </>
  );
};
