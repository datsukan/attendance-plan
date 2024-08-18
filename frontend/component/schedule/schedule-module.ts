import { isAfter, isBefore, isEqual, differenceInCalendarDays } from 'date-fns';

import type { Schedule } from '@/type/schedule';

export function getMasterSchedules(): Schedule[] {
  const sample: Schedule[] = [
    { id: '1', name: '予定1', startDate: new Date(2024, 7, 1), endDate: new Date(2024, 7, 3), styleClassName: 'bg-blue-500 text-white' },
    { id: '2', name: '予定2', startDate: new Date(2024, 7, 2), endDate: new Date(2024, 7, 3), styleClassName: 'bg-orange-500 text-white' },
    { id: '8', name: '予定8', startDate: new Date(2024, 7, 4), endDate: new Date(2024, 7, 4), styleClassName: 'bg-pink-500 text-white' },
    { id: '3', name: '予定3', startDate: new Date(2024, 7, 5), endDate: new Date(2024, 7, 7), styleClassName: 'bg-green-500 text-white' },
    { id: '4', name: '予定4', startDate: new Date(2024, 7, 1), endDate: new Date(2024, 7, 1), styleClassName: 'bg-red-500 text-white' },
    { id: '5', name: '予定5', startDate: new Date(2024, 7, 2), endDate: new Date(2024, 7, 2), styleClassName: 'bg-red-500 text-white' },
    { id: '6', name: '予定6', startDate: new Date(2024, 7, 1), endDate: new Date(2024, 7, 23), styleClassName: 'bg-gray-500 text-white' },
    { id: '7', name: '予定7', startDate: new Date(2024, 7, 1), endDate: new Date(2024, 7, 10), styleClassName: 'bg-yellow-500 text-white' },
  ];

  return sample;
}

export function getCustomSchedules(): Schedule[] {
  const sample: Schedule[] = [
    {
      id: '1',
      name: 'プログラミング教育B',
      startDate: new Date(2024, 7, 1),
      endDate: new Date(2024, 7, 1),
      styleClassName: 'bg-white border border-gray-400',
    },
    {
      id: '2',
      name: '経営学入門B',
      startDate: new Date(2024, 7, 2),
      endDate: new Date(2024, 7, 2),
      styleClassName: 'bg-white border border-gray-400',
    },
    {
      id: '3',
      name: '情報基礎B',
      startDate: new Date(2024, 7, 3),
      endDate: new Date(2024, 7, 3),
      styleClassName: 'bg-white border border-gray-400',
    },
    {
      id: '4',
      name: '初級プログラミング',
      startDate: new Date(2024, 7, 3),
      endDate: new Date(2024, 7, 4),
      styleClassName: 'bg-white border border-gray-400',
    },
  ];

  return sample;
}

export function isDisplaySchedule(schedule: Schedule, date: Date): boolean {
  if (isEqual(schedule.startDate, date)) {
    return true;
  }

  if (isEqual(schedule.endDate, date)) {
    return true;
  }

  return isBefore(schedule.startDate, date) && isAfter(schedule.endDate, date);
}

export function isShowItem(index: number, schedule: Schedule, date: Date): boolean {
  if (date < schedule.startDate || date > schedule.endDate) {
    return false;
  }

  if (index > 0 && isBefore(schedule.startDate, date)) {
    return false;
  }

  return true;
}

export function getColStartClassName(index: number): string {
  let className = '';
  switch (index) {
    case 0:
      className = 'col-start-1';
      break;
    case 1:
      className = 'col-start-2';
      break;
    case 2:
      className = 'col-start-3';
      break;
    case 3:
      className = 'col-start-4';
      break;
    case 4:
      className = 'col-start-5';
      break;
    case 5:
      className = 'col-start-6';
      break;
    default:
      className = 'col-start-7';
      break;
  }

  return className;
}

export function getColEndClassName(index: number, schedule: Schedule, dates: Date[]): string {
  let range = 0;
  if (isBefore(dates[0], schedule.startDate)) {
    const sub = differenceInCalendarDays(schedule.endDate, schedule.startDate);
    range = index + sub + 1;
  } else {
    const sub = differenceInCalendarDays(schedule.endDate, dates[0]);
    range = sub + 1;
  }

  let className = '';
  switch (range) {
    case 1:
      className = 'col-end-2';
      break;
    case 2:
      className = 'col-end-3';
      break;
    case 3:
      className = 'col-end-4';
      break;
    case 4:
      className = 'col-end-5';
      break;
    case 5:
      className = 'col-end-6';
      break;
    case 6:
      className = 'col-end-7';
      break;
    default:
      className = 'col-end-8';
      break;
  }

  return className;
}

export function hasDateLabel(schedule: Schedule): boolean {
  const sub = differenceInCalendarDays(schedule.endDate, schedule.startDate);
  return sub > 0;
}
