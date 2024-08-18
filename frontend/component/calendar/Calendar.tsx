'use client';

import { useState, useMemo, useRef, useEffect } from 'react';
import { getYear, getMonth, startOfMonth, format } from 'date-fns';

import { Schedule } from '@/type/schedule';

import { Header } from './Header';
import { CalendarDateItem } from './CalendarDateItem';
import { ScheduleWeeks } from '@/component/schedule/ScheduleWeeks';

import { getDates, getWeeks, addViewMonths, changeYearMonth, changeDatesByMonths, prev, next, reset, dateKey } from './calendar-module';

export const Calender = () => {
  const now = useMemo(() => new Date(), []);

  const [months, setMonths] = useState(1);
  const [dates, setDates] = useState(getDates(startOfMonth(now), months));
  const [weeks, setWeeks] = useState(getWeeks(dates));
  const [masterSchedules, setMasterSchedules] = useState<Schedule[]>([]);
  const [customSchedules, setCustomSchedules] = useState<Schedule[]>([]);
  const [year, setYear] = useState(getYear(now));
  const [month, setMonth] = useState(getMonth(now));
  const [baseDate, setBaseDate] = useState(startOfMonth(now));
  const [weekHeights, setWeekHeights] = useState<number[]>([]);

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

  useEffect(() => {
    setWeeks(getWeeks(dates));
  }, [dates]);

  const changeWeekHeight = (index: number, height: number) => {
    const newHeights = weekHeights;
    newHeights[index] = height;
    setWeekHeights(newHeights);
  };

  const addSchedule = (name: string, startDate: Date, endDate: Date) => {
    console.log(name, format(startDate, 'yyyy-MM-dd HH:mm'), format(endDate, 'yyyy-MM-dd HH:mm'));
    const schedule: Schedule = {
      id: format(new Date(), 'yyyyMMddHHmmss'),
      name,
      startDate,
      endDate,
      styleClassName: 'bg-white border border-gray-400',
    };
    setCustomSchedules([...customSchedules, schedule]);
    console.log(customSchedules);
  };

  return (
    <div className="relative">
      <div className="sticky top-0 bg-white z-10">
        <Header
          year={year}
          month={month + 1}
          prev={() => prev(baseDate, setDates, setYear, setMonth, setBaseDate, setMonths)}
          next={() => next(baseDate, setDates, setYear, setMonth, setBaseDate, setMonths)}
          reset={() => reset(setDates, setYear, setMonth, setBaseDate, setMonths)}
          create={addSchedule}
        />
      </div>
      <div className="border-r border-b">
        {weeks.map((week, i) => {
          return (
            <div key={dateKey(week[0]) + '-frame-week'} className="relative">
              <div
                data-date={dateKey(week[0])}
                className="grid grid-cols-7 calendar-item"
                style={{ minHeight: (weekHeights[i] ?? 0) + 50 }}
              >
                {week.map((date) => {
                  return (
                    <div key={dateKey(date) + '-frame-date'} className="border-t border-l">
                      <CalendarDateItem date={date} />
                    </div>
                  );
                })}
              </div>
              <div className="absolute top-10 bottom-0 left-0 right-0">
                <ScheduleWeeks
                  index={i}
                  week={week}
                  changeHeight={changeWeekHeight}
                  masterSchedules={masterSchedules}
                  customSchedules={customSchedules}
                />
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};
