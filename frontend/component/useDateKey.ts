import { format, parse } from 'date-fns';

export const dateKeyFormat = 'yyyy-MM-dd';

export const useDateKey = () => {
  const dateToKey = (date: Date) => {
    return format(date, dateKeyFormat);
  };

  const keyToDate = (key: string) => {
    return parse(key, dateKeyFormat, new Date());
  };

  return {
    dateToKey,
    keyToDate,
  };
};
