import { RefObject } from 'react';

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
  startOfDay,
} from 'date-fns';
import type { StartOfWeekOptions, EndOfWeekOptions } from 'date-fns';

const dateKeyFormat = 'yyyy-MM-dd';

export function getDates(date: Date, months: number): Date[] {
  const options = { weekStartsOn: 1 };
  const start = startOfWeek(date, options as StartOfWeekOptions);
  const end = endOfWeek(addMonths(endOfMonth(date), months), options as EndOfWeekOptions);

  return eachDayOfInterval({ start, end });
}

export function addViewMonths(months: number, setMonths: (months: number) => void): () => void {
  const bodyHeight = document.body.scrollHeight;
  const windowHeight = window.innerHeight;
  const bottomPoint = bodyHeight - windowHeight - 200;

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
}

export function changeYearMonth(
  calendarRef: RefObject<HTMLDivElement>,
  firstDate: Date,
  baseDate: Date,
  setYear: (year: number) => void,
  setMonth: (month: number) => void
): () => void {
  const handleScroll = () => {
    if (!calendarRef.current) {
      return;
    }

    const items = calendarRef.current.querySelectorAll('.calendar-item');
    for (let item of items) {
      const rect = item.getBoundingClientRect();

      if (rect.top >= 0) {
        let date = parse(item.getAttribute('data-date') as string, dateKeyFormat, new Date());
        if (isEqual(date, firstDate) && (getYear(date) !== getYear(baseDate) || getMonth(date) !== getMonth(baseDate))) {
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
}

export function changeDatesByMonths(
  baseDate: Date,
  months: number,
  setDates: (dates: Date[]) => void,
  getDates: (date: Date, months: number) => Date[]
): void {
  setDates(getDates(startOfMonth(baseDate), months));
}

export function initCalendar(
  date: Date,
  setDates: (dates: Date[]) => void,
  setYear: (year: number) => void,
  setMonth: (month: number) => void,
  setBaseDate: (date: Date) => void,
  setMonths: (months: number) => void
): void {
  scrollTo(0, 0);

  setMonths(1);
  setBaseDate(date);
  setDates(getDates(date, 1));
  setYear(getYear(date));
  setMonth(getMonth(date));
}

export function prev(
  baseDate: Date,
  setDates: (dates: Date[]) => void,
  setYear: (year: number) => void,
  setMonth: (month: number) => void,
  setBaseDate: (date: Date) => void,
  setMonths: (months: number) => void
): void {
  const date = subMonths(startOfDay(baseDate), 1);
  initCalendar(date, setDates, setYear, setMonth, setBaseDate, setMonths);
}

export function next(
  baseDate: Date,
  setDates: (dates: Date[]) => void,
  setYear: (year: number) => void,
  setMonth: (month: number) => void,
  setBaseDate: (date: Date) => void,
  setMonths: (months: number) => void
): void {
  const date = addMonths(startOfDay(baseDate), 1);
  initCalendar(date, setDates, setYear, setMonth, setBaseDate, setMonths);
}

export function reset(
  setDates: (dates: Date[]) => void,
  setYear: (year: number) => void,
  setMonth: (month: number) => void,
  setBaseDate: (date: Date) => void,
  setMonths: (months: number) => void
) {
  const date = startOfMonth(new Date());
  initCalendar(date, setDates, setYear, setMonth, setBaseDate, setMonths);
}

export function dateKey(date: Date): string {
  return format(date, dateKeyFormat);
}
