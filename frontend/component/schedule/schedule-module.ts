import { isAfter, isBefore, isEqual, differenceInCalendarDays, format } from 'date-fns';

import { Type } from '@/type';

export function isDisplaySchedule(schedule: Type.Schedule, date: Date): boolean {
  if (isEqual(schedule.startDate, date)) {
    return true;
  }

  if (isEqual(schedule.endDate, date)) {
    return true;
  }

  return isBefore(schedule.startDate, date) && isAfter(schedule.endDate, date);
}

export function isShowItem(index: number, schedule: Type.Schedule, date: Date): boolean {
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

export function getColEndClassName(schedule: Type.Schedule, dates: Date[]): string {
  // Both "schedule starts this week" and "cross-week continuation" cases reduce to
  // the same formula: days from week start to schedule end + 1.
  const range = differenceInCalendarDays(schedule.endDate, dates[0]) + 1;

  switch (range) {
    case 1:
      return 'col-end-2';
    case 2:
      return 'col-end-3';
    case 3:
      return 'col-end-4';
    case 4:
      return 'col-end-5';
    case 5:
      return 'col-end-6';
    case 6:
      return 'col-end-7';
    default:
      return 'col-end-8';
  }
}

export function hasDateLabel(schedule: Type.Schedule): boolean {
  const sub = differenceInCalendarDays(schedule.endDate, schedule.startDate);
  return sub > 0;
}

export function getMasterScheduleTemplates(): { name: string; color: string }[] {
  type template = { name: string; color: string };
  const templates: template[] = [
    { name: '履修登録', color: 'red' },
    { name: '第1回 授業配信', color: 'yellow' },
    { name: '第2回 授業配信', color: 'yellow' },
    { name: '第3回 授業配信', color: 'yellow' },
    { name: '第4回 授業配信', color: 'yellow' },
    { name: '第5回 授業配信', color: 'yellow' },
    { name: '第6回 授業配信', color: 'yellow' },
    { name: '第7回 授業配信', color: 'yellow' },
    { name: '第8回 授業配信', color: 'yellow' },
    { name: '単位認定試験', color: 'green' },
    { name: '単位認定試験 追試', color: 'green' },
    { name: '成績評価', color: 'blue' },
    { name: '成績発表', color: 'blue' },
  ];

  return templates;
}

export function toScheduleTypeName(type: 'master' | 'custom'): string {
  return type === 'master' ? '学事' : '受講';
}
