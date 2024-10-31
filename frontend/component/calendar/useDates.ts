import { useEffect, useState, useRef } from 'react';

import { startOfWeek, endOfWeek, eachDayOfInterval, addMonths, endOfMonth } from 'date-fns';

export const useDates = (targetDate: Date) => {
  const [dates, setDates] = useState<Date[]>([]);
  const [weeks, setWeeks] = useState<Date[][]>([]);
  const [monthCount, setMonthCount] = useState(1);

  const refreshDates = () => {
    const start = startOfWeek(targetDate, { weekStartsOn: 1 });
    const end = endOfWeek(addMonths(endOfMonth(targetDate), monthCount), { weekStartsOn: 1 });

    const ds = eachDayOfInterval({ start, end });
    setDates(ds);
  };

  const refreshWeeks = () => {
    const ws: Date[][] = [];

    let week: Date[] = [];
    for (let date of dates) {
      week.push(date);

      if (week.length === 7) {
        ws.push(week);
        week = [];
      }
    }

    setWeeks(ws);
  };

  const addMonthCount = () => {
    setMonthCount(monthCount + 1);
  };

  useEffect(() => {
    refreshDates();
    refreshWeeks();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [targetDate, monthCount]);

  useEffect(() => {
    refreshWeeks();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [dates]);

  return {
    dates,
    weeks,
    monthCount,
    addMonthCount,
  };
};
