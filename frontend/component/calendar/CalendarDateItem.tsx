import { getDate, getDay, isToday, isFirstDayOfMonth, format } from 'date-fns';

type Props = {
  date: Date;
};

export const CalendarDateItem = ({ date }: Props) => {
  return (
    <div className={`flex h-full min-h-40 flex-col items-center py-1 ${isHoliday(date) ? 'bg-gray-50' : ''}`}>
      <div className={`flex h-8 w-8 items-center justify-center ${isToday(date) ? 'rounded-full bg-blue-500 text-white' : ''}`}>
        <span className="text-sm">{isFirstDayOfMonth(date) ? format(date, 'M/d') : getDate(date)}</span>
      </div>
    </div>
  );
};

function isHoliday(date: Date) {
  return getDay(date) === 0 || getDay(date) === 6;
}
