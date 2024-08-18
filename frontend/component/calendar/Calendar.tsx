'use client';

import { useState, useMemo, useRef, useEffect } from 'react';
import { getYear, getMonth, startOfMonth } from 'date-fns';

import { Header } from './Header';
import { CalendarDateItem } from './CalendarDateItem';

import { getDates, addViewMonths, changeYearMonth, changeDatesByMonths, prev, next, reset, dateKey } from './calendar-module';

export const Calender = () => {
  const now = useMemo(() => new Date(), []);

  const [months, setMonths] = useState(1);
  const [dates, setDates] = useState(getDates(startOfMonth(now), months));
  const [year, setYear] = useState(getYear(now));
  const [month, setMonth] = useState(getMonth(now));
  const [baseDate, setBaseDate] = useState(startOfMonth(now));

  const calendarRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    scrollTo(0, 0);
  }, [baseDate]);

  useEffect(() => {
    return addViewMonths(months, setMonths);
  }, [months]);

  useEffect(() => {
    return changeYearMonth(calendarRef, dates[0], baseDate, setYear, setMonth);
  }, [dates, baseDate]);

  useEffect(() => {
    changeDatesByMonths(baseDate, months, setDates, getDates);
  }, [baseDate, months]);

  return (
    <div className="relative">
      <div className="sticky top-0 bg-white">
        <Header
          year={year}
          month={month}
          prev={() => prev(baseDate, setDates, setYear, setMonth, setBaseDate, setMonths)}
          next={() => next(baseDate, setDates, setYear, setMonth, setBaseDate, setMonths)}
          reset={() => reset(setDates, setYear, setMonth, setBaseDate, setMonths)}
        />
      </div>
      <div className="grid grid-cols-7 border-r border-b" ref={calendarRef}>
        {dates.map((date) => {
          return (
            <div key={dateKey(date)} data-date={dateKey(date)} className="border-t border-l calendar-item">
              <CalendarDateItem date={date} />
            </div>
          );
        })}
      </div>
    </div>
  );
};
