import { useState, useEffect } from 'react';
import type { RefObject } from 'react';

import { getYear, getMonth, startOfMonth, endOfWeek, addMonths, subMonths, parse } from 'date-fns';

const dateKeyFormat = 'yyyy-MM-dd';

export const useTargetDate = (initDate: Date, calendarRef: RefObject<HTMLDivElement>) => {
  const [date, setDate] = useState(startOfMonth(initDate));
  const [year, setYear] = useState(getYear(date));
  const [month, setMonth] = useState(getMonth(date) + 1);

  const setAll = (d: Date) => {
    setDate(d);
    setYear(getYear(d));
    setMonth(getMonth(d) + 1);
  };

  const incrementMonth = () => {
    const d = addMonths(date, 1);
    setAll(d);
  };

  const decrementMonth = () => {
    const d = subMonths(date, 1);
    setAll(d);
  };

  const resetMonth = () => {
    const d = startOfMonth(new Date());
    setAll(d);
  };

  useEffect(() => {
    const handleScroll = () => {
      if (!calendarRef.current) {
        return;
      }

      const items = calendarRef.current.querySelectorAll('.calendar-item');
      for (let item of items) {
        const rect = item.getBoundingClientRect();

        if (rect.top < 0) {
          continue;
        }

        const d = parse(item.getAttribute('data-date') as string, dateKeyFormat, new Date());
        const after = endOfWeek(d, { weekStartsOn: 1 });
        setYear(getYear(after));
        setMonth(getMonth(after) + 1);

        break;
      }
    };

    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, [calendarRef]);

  return {
    targetDate: date,
    targetYear: year,
    targetMonth: month,
    incrementMonth,
    decrementMonth,
    resetMonth,
  };
};
