'use client';

import { useState, useMemo, useRef, useEffect } from 'react';
import { getYear, getMonth, startOfMonth } from 'date-fns';

import { Schedule } from '@/type/schedule';

import { Header } from './Header';
import { CalenderDates } from './CalenderDates';

import { getDates, getWeeks, addViewMonths, changeYearMonth, changeDatesByMonths, prev, next, reset } from './calendar-module';
import { CreateSchedule } from '@/model/create-schedule';
import { EditSchedule } from '@/model/edit-schedule';

import { useInitPagePosition } from '@/hook/useInitPagePosition';

export const Calender = () => {
  const now = useMemo(() => new Date(), []);

  const [months, setMonths] = useState(1);
  const [dates, setDates] = useState(getDates(startOfMonth(now), months));
  const [weeks, setWeeks] = useState(getWeeks(dates));
  const [schedules, setSchedules] = useState<Schedule[]>([]);
  const [year, setYear] = useState(getYear(now));
  const [month, setMonth] = useState(getMonth(now));
  const [baseDate, setBaseDate] = useState(startOfMonth(now));

  const calendarRef = useRef<HTMLDivElement>(null);

  useInitPagePosition(baseDate);

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
      <CalenderDates
        weeks={weeks}
        schedules={schedules}
        setSchedules={setSchedules}
        removeSchedule={removeSchedule}
        saveSchedule={saveSchedule}
        changeScheduleColor={changeScheduleColor}
      />
    </div>
  );
};
