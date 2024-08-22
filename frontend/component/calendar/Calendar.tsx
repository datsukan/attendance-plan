'use client';

import { useState, useMemo, useRef, useEffect } from 'react';
import { getYear, getMonth, startOfMonth, format } from 'date-fns';

import { Schedule } from '@/type/schedule';

import { Header } from './Header';
import { CalendarDateItem } from './CalendarDateItem';
import { ScheduleWeeks } from '@/component/schedule/ScheduleWeeks';

import { getDates, getWeeks, addViewMonths, changeYearMonth, changeDatesByMonths, prev, next, reset, dateKey } from './calendar-module';
import { CreateSchedule } from '@/model/create-schedule';
import { EditSchedule } from '@/model/edit-schedule';

export const Calender = () => {
  const now = useMemo(() => new Date(), []);

  const [months, setMonths] = useState(1);
  const [dates, setDates] = useState(getDates(startOfMonth(now), months));
  const [weeks, setWeeks] = useState(getWeeks(dates));
  const [schedules, setSchedules] = useState<Schedule[]>([]);
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
    if (height === 0) {
      return;
    }
    const newHeights = weekHeights;
    newHeights[index] = height;
    setWeekHeights(newHeights);
  };

  useEffect(() => {
    setWeekHeights([...weekHeights]);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [schedules]);

  const addSchedule = (createSchedule: CreateSchedule) => {
    const afterSchedules: Schedule[] = [...schedules];
    if (createSchedule.getHasBulk()) {
      for (let number = createSchedule.getBulkFrom(); number <= createSchedule.getBulkTo(); number++) {
        afterSchedules.push({
          id: createSchedule.getId() + '-' + number,
          name: `第${number}回 ${createSchedule.getName()}`,
          startDate: createSchedule.getStartDate(),
          endDate: createSchedule.getEndDate(),
          color: createSchedule.getColor(),
          type: createSchedule.getType(),
        });
      }
    } else {
      afterSchedules.push({
        id: createSchedule.getId(),
        name: createSchedule.getName(),
        startDate: createSchedule.getStartDate(),
        endDate: createSchedule.getEndDate(),
        color: createSchedule.getColor(),
        type: createSchedule.getType(),
      });
    }

    setSchedules(afterSchedules);
  };

  const removeSchedule = (id: string) => {
    setSchedules(schedules.filter((schedule) => schedule.id !== id));
  };

  const saveSchedule = (editSchedule: EditSchedule) => {
    const set = (schedule: Schedule) => {
      if (schedule.id !== editSchedule.getId()) {
        return schedule;
      }

      schedule.name = editSchedule.getName();
      schedule.startDate = editSchedule.getStartDate();
      schedule.endDate = editSchedule.getEndDate();
      schedule.color = editSchedule.getColor();
      schedule.type = editSchedule.getType();
      return schedule;
    };

    setSchedules(schedules.map(set));
  };

  const changeScheduleColor = (id: string, color: string) => {
    const set = (schedule: Schedule) => {
      if (schedule.id !== id) {
        return schedule;
      }

      schedule.color = color;
      return schedule;
    };

    setSchedules(schedules.map(set));
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
          const minHeight = weekHeights[i] ? weekHeights[i] + 50 : 0;
          return (
            <div key={dateKey(week[0]) + '-frame-week'} className="relative">
              <div data-date={dateKey(week[0])} className="grid grid-cols-7 calendar-item" style={{ minHeight: minHeight }}>
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
                  masterSchedules={schedules.filter((schedule) => schedule.type === 'master')}
                  customSchedules={schedules.filter((schedule) => schedule.type === 'custom')}
                  removeSchedule={removeSchedule}
                  saveSchedule={saveSchedule}
                  changeScheduleColor={changeScheduleColor}
                />
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};
