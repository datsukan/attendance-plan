import { isAfter, isBefore, isEqual, differenceInCalendarDays } from 'date-fns';

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

const COL_STARTS = [
  'col-start-1', 'col-start-2', 'col-start-3', 'col-start-4',
  'col-start-5', 'col-start-6', 'col-start-7',
];

export function getColStartClassName(index: number): string {
  return COL_STARTS[index] ?? 'col-start-7';
}

// 範囲 1〜6 は対応する col-end クラス、範囲外（週をまたぐ場合など）は col-end-8 を返す
const COL_ENDS = [
  'col-end-2', 'col-end-3', 'col-end-4', 'col-end-5', 'col-end-6', 'col-end-7',
];

export function getColEndClassName(schedule: Type.Schedule, dates: Date[]): string {
  // 週の先頭から終了日までの日数 + 1 = 必要な列数
  const range = differenceInCalendarDays(schedule.endDate, dates[0]) + 1;
  return COL_ENDS[range - 1] ?? 'col-end-8';
}

export function hasDateLabel(schedule: Type.Schedule): boolean {
  const sub = differenceInCalendarDays(schedule.endDate, schedule.startDate);
  return sub > 0;
}

export function getMasterScheduleTemplates(): { name: string; color: string }[] {
  return [
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
}

export function toScheduleTypeName(type: 'master' | 'custom'): string {
  return type === 'master' ? '学事' : '受講';
}
