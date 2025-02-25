'use client';

import { useMemo, useRef, useEffect } from 'react';

import { Header } from './Header';
import { CalenderDates } from './CalenderDates';

import { usePagePosition } from './usePagePosition';
import { useTargetDate } from './useTargetDate';
import { useDates } from './useDates';
import { useSchedule } from '@/provider/ScheduleProvider';

export const Calender = () => {
  const now = useMemo(() => new Date(), []);
  const calendarRef = useRef<HTMLDivElement>(null);

  const { targetDate, targetYear, targetMonth, incrementMonth, decrementMonth, resetMonth } = useTargetDate(now, calendarRef);
  const { weeks, monthCount, addMonthCount } = useDates(targetDate);
  const { initPagePosition, execWhenPageBottom } = usePagePosition();
  const { addSchedule } = useSchedule();

  useEffect(() => {
    initPagePosition();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  useEffect(() => {
    execWhenPageBottom(addMonthCount);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [monthCount]);

  const prev = () => {
    decrementMonth();
    initPagePosition();
  };

  const next = () => {
    incrementMonth();
    initPagePosition();
  };

  const reset = () => {
    resetMonth();
    initPagePosition();
  };

  return (
    <div className="relative" ref={calendarRef}>
      <div className="sticky top-0 z-10 bg-white">
        <Header year={targetYear} month={targetMonth} prev={prev} next={next} reset={reset} create={addSchedule} />
      </div>
      <CalenderDates weeks={weeks} />
    </div>
  );
};
