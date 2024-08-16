import { getDate, getDay, isToday, isFirstDayOfMonth, format } from 'date-fns';

type Props = {
  date: Date;
};

export const CalendarDateItem = ({ date }: Props) => {
  return (
    <div className={`min-h-40 py-1 flex flex-col items-center ${isHoliday(date) ? 'bg-gray-50' : ''}`}>
      <div className={`w-8 h-8 flex justify-center items-center ${isToday(date) ? 'rounded-full bg-blue-500 text-white' : ''}`}>
        <span className="text-sm">{isFirstDayOfMonth(date) ? format(date, 'M/d') : getDate(date)}</span>
      </div>
    </div>
  );
};

function isHoliday(date: Date) {
  return getDay(date) === 0 || getDay(date) === 6;
}
