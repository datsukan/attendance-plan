'use client';

import { useState, useMemo, useRef, useEffect } from 'react';
import {
  getYear,
  getMonth,
  startOfWeek,
  endOfWeek,
  startOfMonth,
  endOfMonth,
  eachDayOfInterval,
  addMonths,
  subMonths,
  isEqual,
  format,
  parse,
} from 'date-fns';
import type { StartOfWeekOptions, EndOfWeekOptions } from 'date-fns';
import { ChevronLeftIcon, ChevronRightIcon } from '@heroicons/react/24/outline';

import { NavButton } from './NavButton';
import { DayOfWeeks } from './DayOfWeeks';
import { CalendarDateItem } from './CalendarDateItem';

export const Calender = () => {
  const now = useMemo(() => new Date(), []);

  const [months, setMonths] = useState(1);
  const [dates, setDates] = useState(getDates(startOfMonth(now), months));
  const [year, setYear] = useState(getYear(now));
  const [month, setMonth] = useState(getMonth(now));
  const [baseDate, setBaseDate] = useState(startOfMonth(now));

  const calendarRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const bodyHeight = document.body.clientHeight;
    const windowHeight = window.innerHeight;
    const bottomPoint = bodyHeight - windowHeight;

    const handleScroll = () => {
      const currentPos = window.scrollY;
      if (bottomPoint <= currentPos) {
        setMonths(months + 1);
      }
    };

    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, [months]);

  useEffect(() => {
    scrollTo(0, 0);

    const handleScroll = () => {
      if (!calendarRef.current) {
        return;
      }

      const items = calendarRef.current.querySelectorAll('.calendar-item');
      for (let item of items) {
        const rect = item.getBoundingClientRect();

        if (rect.top >= 0) {
          let date = parse(item.getAttribute('data-date') as string, 'yyyy-MM-dd', new Date());
          if (isEqual(date, dates[0]) && (getYear(date) !== getYear(baseDate) || getMonth(date) !== getMonth(baseDate))) {
            console.log(format(date, 'yyyy-MM-dd'));
            date = addMonths(date, 1);
            setYear(getYear(date));
            setMonth(getMonth(date));
          } else {
            setYear(getYear(date));
            setMonth(getMonth(date));
          }
          break;
        }
      }
    };

    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [baseDate]);

  useEffect(() => {
    setDates(getDates(startOfMonth(baseDate), months));
  }, [baseDate, months]);

  const initCalendar = (date: Date) => {
    scrollTo(0, 0);

    setBaseDate(date);
    setDates(getDates(date, 1));
    setYear(getYear(date));
    setMonth(getMonth(date));
  };

  const prev = () => {
    setMonths(1);

    const date = subMonths(new Date(year, month, 1), 1);
    initCalendar(date);
  };

  const next = () => {
    setMonths(1);

    const date = addMonths(new Date(year, month, 1), 1);
    initCalendar(date);
  };

  const reset = () => {
    setMonths(1);

    const date = startOfMonth(now);
    initCalendar(date);
  };

  return (
    <div className="relative">
      <div className="sticky top-0 bg-white">
        <div className="py-6 flex gap-6 items-center justify-between">
          <h2 className="text-lg">
            {year}年 {month + 1}月
          </h2>
          <div className="flex gap-2">
            <NavButton onClick={() => prev()}>
              <ChevronLeftIcon className="size-5" />
            </NavButton>
            <NavButton onClick={() => reset()}>今日</NavButton>
            <NavButton onClick={() => next()}>
              <ChevronRightIcon className="size-5" />
            </NavButton>
          </div>
        </div>
        <DayOfWeeks />
      </div>
      <div className="grid grid-cols-7 border-r border-b" ref={calendarRef}>
        {dates.map((date) => {
          const dateString = format(date, 'yyyy-MM-dd');

          return (
            <div key={dateString} data-date={dateString} className="border-t border-l calendar-item">
              <CalendarDateItem date={date} />
            </div>
          );
        })}
      </div>
    </div>
  );
};

const getDates = (date: Date, months: number): Date[] => {
  const options = { weekStartsOn: 1 };
  const start = startOfWeek(date, options as StartOfWeekOptions);
  const end = endOfWeek(addMonths(endOfMonth(date), months), options as EndOfWeekOptions);

  return eachDayOfInterval({ start, end });
};
