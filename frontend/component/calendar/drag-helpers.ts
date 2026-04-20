import { differenceInDays, addDays, isSameDay } from 'date-fns';

import { Type } from '@/type';

/**
 * スケジュールを新しい開始日へシフトする。
 * 期間（duration）を維持したまま startDate / endDate を再計算して返す。
 */
export function shiftSchedule(
  schedule: Type.Schedule,
  newStartDate: Date
): Type.Schedule {
  const duration = differenceInDays(schedule.endDate, schedule.startDate);
  return {
    ...schedule,
    startDate: newStartDate,
    endDate: addDays(newStartDate, duration),
  };
}

/**
 * スケジュールの位置（日付・タイプ）が実際に変わったかを判定する。
 */
export function hasScheduleMoved(
  before: Type.Schedule,
  after: Type.Schedule
): boolean {
  return !isSameDay(before.startDate, after.startDate) || before.type !== after.type;
}

/**
 * バルクドラッグ終了時に、全選択アイテムの移動後スケジュール配列を構築する。
 *
 * - primaryId のスケジュールは DragOver で既に startDate が更新済みのため、
 *   currentPrimaryStart をそのまま使う（origins を基準にすると二重シフトになる）。
 * - それ以外のアイテムは originStart + dateDelta で新しい開始日を決める。
 * - 全アイテムのタイプを targetType に統一する。
 */
export function buildBulkMoves(
  schedules: Type.Schedule[],
  primaryId: string,
  origins: Map<string, Date>,
  targetType: Type.ScheduleType,
  currentPrimaryStart: Date
): Type.Schedule[] {
  const primaryOrigin = origins.get(primaryId);
  if (!primaryOrigin) return [];

  const dateDelta = differenceInDays(currentPrimaryStart, primaryOrigin);

  return schedules
    .filter((s) => origins.has(s.id))
    .map((s) => {
      const origin = origins.get(s.id);
      if (!origin) return null;
      const newStart = s.id === primaryId ? currentPrimaryStart : addDays(origin, dateDelta);
      return {
        ...shiftSchedule(s, newStart),
        type: targetType,
      };
    })
    .filter((s): s is Type.Schedule => s !== null);
}

/**
 * バルクドラッグが実際に移動を伴ったかを判定する。
 */
export function hasBulkMoved(
  primaryId: string,
  origins: Map<string, Date>,
  currentPrimaryStart: Date,
  currentPrimaryType: Type.ScheduleType,
  snapshot: Type.Schedule[]
): boolean {
  const primaryOrigin = origins.get(primaryId);
  if (!primaryOrigin) return false;

  const dateDelta = differenceInDays(currentPrimaryStart, primaryOrigin);
  if (dateDelta !== 0) return true;

  const primarySnapshot = snapshot.find((s) => s.id === primaryId);
  return !!primarySnapshot && currentPrimaryType !== primarySnapshot.type;
}
