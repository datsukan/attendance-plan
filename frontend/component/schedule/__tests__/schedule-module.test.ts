import { describe, it, expect } from 'vitest';
import { isShowItem, isDisplaySchedule } from '../schedule-module';
import type { Type } from '@/type';

const d = (str: string) => new Date(str);

const makeSchedule = (overrides: Partial<Type.Schedule>): Type.Schedule => ({
  id: 'test',
  name: 'Test',
  startDate: d('2024-04-01'),
  endDate: d('2024-04-01'),
  color: 'blue',
  type: 'master',
  order: 1,
  ...overrides,
});

describe('isDisplaySchedule', () => {
  it('startDate と一致する日付で true を返す', () => {
    const schedule = makeSchedule({ startDate: d('2024-04-01'), endDate: d('2024-04-01') });
    expect(isDisplaySchedule(schedule, d('2024-04-01'))).toBe(true);
  });

  it('endDate と一致する日付で true を返す', () => {
    const schedule = makeSchedule({ startDate: d('2024-04-01'), endDate: d('2024-04-03') });
    expect(isDisplaySchedule(schedule, d('2024-04-03'))).toBe(true);
  });

  it('範囲内の中間日付で true を返す', () => {
    const schedule = makeSchedule({ startDate: d('2024-04-01'), endDate: d('2024-04-03') });
    expect(isDisplaySchedule(schedule, d('2024-04-02'))).toBe(true);
  });

  it('範囲外の日付で false を返す', () => {
    const schedule = makeSchedule({ startDate: d('2024-04-01'), endDate: d('2024-04-03') });
    expect(isDisplaySchedule(schedule, d('2024-04-04'))).toBe(false);
  });
});

describe('isShowItem', () => {
  describe('単一日付スケジュール', () => {
    it('index=0 で当日なら true を返す', () => {
      const schedule = makeSchedule({ startDate: d('2024-04-01'), endDate: d('2024-04-01') });
      expect(isShowItem(0, schedule, d('2024-04-01'))).toBe(true);
    });

    it('index=3 で当日なら true を返す', () => {
      const schedule = makeSchedule({ startDate: d('2024-04-04'), endDate: d('2024-04-04') });
      expect(isShowItem(3, schedule, d('2024-04-04'))).toBe(true);
    });

    it('範囲外の日付で false を返す', () => {
      const schedule = makeSchedule({ startDate: d('2024-04-01'), endDate: d('2024-04-01') });
      expect(isShowItem(0, schedule, d('2024-04-02'))).toBe(false);
    });
  });

  describe('範囲日付スケジュール（週内）', () => {
    it('startDate (index=0) で true を返す', () => {
      const schedule = makeSchedule({ startDate: d('2024-04-01'), endDate: d('2024-04-03') });
      expect(isShowItem(0, schedule, d('2024-04-01'))).toBe(true);
    });

    it('中間日 (index>0, startDate < date) で false を返す', () => {
      const schedule = makeSchedule({ startDate: d('2024-04-01'), endDate: d('2024-04-03') });
      expect(isShowItem(1, schedule, d('2024-04-02'))).toBe(false);
    });

    it('endDate (index>0, startDate < date) で false を返す', () => {
      const schedule = makeSchedule({ startDate: d('2024-04-01'), endDate: d('2024-04-03') });
      expect(isShowItem(2, schedule, d('2024-04-03'))).toBe(false);
    });
  });

  describe('週跨ぎスケジュール（継続表示）', () => {
    // スケジュールが前週から継続する場合、次週先頭（index=0）でも表示される。
    // これは仕様上の意図的な動作で、ScheduleWeek.tsx 側で
    // isBefore(schedule.startDate, date) を使い ScheduleItem として描画することで
    // useSortable の二重登録を防ぐ。

    it('次週先頭 (index=0, startDate < date) で true を返す（継続表示の仕様）', () => {
      // Sun Apr 7 → Tue Apr 9 のスケジュールが次週 Mon Apr 8 (index=0) でも表示される
      const schedule = makeSchedule({ startDate: d('2024-04-07'), endDate: d('2024-04-09') });
      expect(isShowItem(0, schedule, d('2024-04-08'))).toBe(true);
    });

    it('次週の index>0 では false を返す', () => {
      const schedule = makeSchedule({ startDate: d('2024-04-07'), endDate: d('2024-04-09') });
      expect(isShowItem(1, schedule, d('2024-04-09'))).toBe(false);
    });

    it('週跨ぎ継続日は isDisplaySchedule でも範囲内と判定される', () => {
      const schedule = makeSchedule({ startDate: d('2024-04-07'), endDate: d('2024-04-09') });
      expect(isDisplaySchedule(schedule, d('2024-04-08'))).toBe(true);
    });

    it('スケジュール終了後の日付は isDisplaySchedule で false を返す', () => {
      const schedule = makeSchedule({ startDate: d('2024-04-07'), endDate: d('2024-04-09') });
      expect(isDisplaySchedule(schedule, d('2024-04-10'))).toBe(false);
    });
  });
});
