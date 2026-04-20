import { describe, it, expect } from 'vitest';
import { addDays, format } from 'date-fns';

import { shiftSchedule, hasScheduleMoved, buildBulkMoves, hasBulkMoved } from '../drag-helpers';
import type { Type } from '@/type';

const d = (str: string) => new Date(str);

const makeSchedule = (overrides: Partial<Type.Schedule> = {}): Type.Schedule => ({
  id: 'schedule-1',
  name: 'テスト',
  startDate: d('2024-04-01'),
  endDate: d('2024-04-01'),
  color: 'blue',
  type: 'master',
  order: 1,
  ...overrides,
});

// ─── shiftSchedule ───────────────────────────────────────────────────────────

describe('shiftSchedule', () => {
  it('1日スケジュールを3日先にシフトする', () => {
    const s = makeSchedule({
      startDate: d('2024-04-01'),
      endDate: d('2024-04-01'),
    });
    const result = shiftSchedule(s, d('2024-04-04'));
    expect(format(result.startDate, 'yyyy-MM-dd')).toBe('2024-04-04');
    expect(format(result.endDate, 'yyyy-MM-dd')).toBe('2024-04-04');
  });

  it('複数日スケジュールのdurationを保持してシフトする', () => {
    const s = makeSchedule({
      startDate: d('2024-04-01'),
      endDate: d('2024-04-05'), // 4泊5日
    });
    const result = shiftSchedule(s, d('2024-04-10'));
    expect(format(result.startDate, 'yyyy-MM-dd')).toBe('2024-04-10');
    expect(format(result.endDate, 'yyyy-MM-dd')).toBe('2024-04-14'); // duration=4を維持
  });

  it('元のスケジュールを変更しない（immutable）', () => {
    const s = makeSchedule({ startDate: d('2024-04-01'), endDate: d('2024-04-03') });
    const originalStart = s.startDate;
    shiftSchedule(s, d('2024-05-01'));
    expect(s.startDate).toBe(originalStart);
  });

  it('その他のフィールドを維持する', () => {
    const s = makeSchedule({ id: 'abc', name: 'テスト名', color: 'red', type: 'custom', order: 5 });
    const result = shiftSchedule(s, d('2024-06-01'));
    expect(result.id).toBe('abc');
    expect(result.name).toBe('テスト名');
    expect(result.color).toBe('red');
    expect(result.type).toBe('custom');
    expect(result.order).toBe(5);
  });
});

// ─── hasScheduleMoved ────────────────────────────────────────────────────────

describe('hasScheduleMoved', () => {
  it('同じ日付・同じタイプなら false', () => {
    const s = makeSchedule({ startDate: d('2024-04-01'), type: 'master' });
    expect(hasScheduleMoved(s, { ...s })).toBe(false);
  });

  it('日付が変われば true', () => {
    const before = makeSchedule({ startDate: d('2024-04-01') });
    const after = { ...before, startDate: d('2024-04-02') };
    expect(hasScheduleMoved(before, after)).toBe(true);
  });

  it('タイプが変われば true', () => {
    const before = makeSchedule({ type: 'master' });
    const after: Type.Schedule = { ...before, type: 'custom' };
    expect(hasScheduleMoved(before, after)).toBe(true);
  });

  it('日付とタイプが両方変わっても true', () => {
    const before = makeSchedule({ startDate: d('2024-04-01'), type: 'master' });
    const after: Type.Schedule = { ...before, startDate: d('2024-05-01'), type: 'custom' };
    expect(hasScheduleMoved(before, after)).toBe(true);
  });
});

// ─── buildBulkMoves ──────────────────────────────────────────────────────────

describe('buildBulkMoves', () => {
  const primary = makeSchedule({
    id: 'p1',
    startDate: d('2024-04-05'), // DragOver 後の現在位置
    endDate: d('2024-04-05'),
    type: 'custom',
  });
  const secondary = makeSchedule({
    id: 's1',
    startDate: d('2024-04-03'), // DragOver 中は更新されていない
    endDate: d('2024-04-04'),
    type: 'master',
  });

  const origins = new Map<string, Date>([
    ['p1', d('2024-04-01')], // primary の元の日付
    ['s1', d('2024-04-03')], // secondary の元の日付
  ]);

  it('primaryは currentPrimaryStart をそのまま使う', () => {
    const result = buildBulkMoves([primary, secondary], 'p1', origins, 'custom', d('2024-04-05'));
    const p = result.find((s) => s.id === 'p1')!;
    expect(format(p.startDate, 'yyyy-MM-dd')).toBe('2024-04-05');
  });

  it('secondaryはdateDelta（+4日）を加算する', () => {
    // dateDelta = 2024-04-05 - 2024-04-01 = 4
    const result = buildBulkMoves([primary, secondary], 'p1', origins, 'custom', d('2024-04-05'));
    const s = result.find((s) => s.id === 's1')!;
    expect(format(s.startDate, 'yyyy-MM-dd')).toBe('2024-04-07'); // 04-03 + 4 = 04-07
  });

  it('全アイテムのタイプをtargetTypeに統一する', () => {
    const result = buildBulkMoves([primary, secondary], 'p1', origins, 'custom', d('2024-04-05'));
    expect(result.every((s) => s.type === 'custom')).toBe(true);
  });

  it('originsに存在しないスケジュールは除外する', () => {
    const outsider = makeSchedule({ id: 'out', startDate: d('2024-04-10') });
    const result = buildBulkMoves([primary, secondary, outsider], 'p1', origins, 'master', d('2024-04-05'));
    expect(result.find((s) => s.id === 'out')).toBeUndefined();
  });

  it('secondaryのdurationを維持する（2日スパンのスケジュール）', () => {
    const result = buildBulkMoves([primary, secondary], 'p1', origins, 'master', d('2024-04-05'));
    const s = result.find((s) => s.id === 's1')!;
    const duration = (s.endDate.getTime() - s.startDate.getTime()) / (1000 * 60 * 60 * 24);
    expect(duration).toBe(1); // 元のduration: 04-04 - 04-03 = 1
  });
});

// ─── hasBulkMoved ────────────────────────────────────────────────────────────

describe('hasBulkMoved', () => {
  const snapshot = [makeSchedule({ id: 'p1', startDate: d('2024-04-01'), type: 'master' })];
  const origins = new Map<string, Date>([['p1', d('2024-04-01')]]);

  it('日付も変わらず、タイプも変わらなければ false', () => {
    expect(hasBulkMoved('p1', origins, d('2024-04-01'), 'master', snapshot)).toBe(false);
  });

  it('日付が変われば true', () => {
    expect(hasBulkMoved('p1', origins, d('2024-04-05'), 'master', snapshot)).toBe(true);
  });

  it('タイプが変われば true', () => {
    expect(hasBulkMoved('p1', origins, d('2024-04-01'), 'custom', snapshot)).toBe(true);
  });

  it('primaryId が origins に存在しなければ false', () => {
    expect(hasBulkMoved('unknown', origins, d('2024-04-05'), 'master', snapshot)).toBe(false);
  });
});
