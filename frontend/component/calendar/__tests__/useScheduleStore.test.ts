import { describe, it, expect, vi, beforeEach } from 'vitest';
import { renderHook, act } from '@testing-library/react';
import { format } from 'date-fns';

import { useScheduleStore } from '../useScheduleStore';
import type { Type } from '@/type';

// ─── useSchedule モック ───────────────────────────────────────────────────────

const mockSetSchedulesByType = vi.fn();
const mockSetSchedulesByTypeFunctional = vi.fn();

let mockMasterSchedules: Type.ScheduleDateItem[] = [];
let mockCustomSchedules: Type.ScheduleDateItem[] = [];

vi.mock('@/provider/ScheduleProvider', () => ({
  useSchedule: () => ({
    masterSchedules: mockMasterSchedules,
    customSchedules: mockCustomSchedules,
    setSchedulesByType: mockSetSchedulesByType,
    setSchedulesByTypeFunctional: mockSetSchedulesByTypeFunctional,
  }),
}));

// ─── テストヘルパー ────────────────────────────────────────────────────────────

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

const makeDateItem = (
  date: string,
  type: Type.ScheduleType,
  schedules: Type.Schedule[]
): Type.ScheduleDateItem => ({ date, type, schedules });

// ─── テスト ───────────────────────────────────────────────────────────────────

describe('useScheduleStore', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    mockMasterSchedules = [];
    mockCustomSchedules = [];
  });

  // ─── findById ──────────────────────────────────────────────────────────────

  describe('findById', () => {
    it('masterSchedules からIDで検索できる', () => {
      const s = makeSchedule({ id: 'abc' });
      mockMasterSchedules = [makeDateItem('2024-04-01', 'master', [s])];

      const { result } = renderHook(() => useScheduleStore());
      expect(result.current.findById('abc')?.id).toBe('abc');
    });

    it('customSchedules からIDで検索できる', () => {
      const s = makeSchedule({ id: 'xyz', type: 'custom' });
      mockCustomSchedules = [makeDateItem('2024-04-01', 'custom', [s])];

      const { result } = renderHook(() => useScheduleStore());
      expect(result.current.findById('xyz')?.id).toBe('xyz');
    });

    it('存在しないIDは undefined を返す', () => {
      const { result } = renderHook(() => useScheduleStore());
      expect(result.current.findById('not-exists')).toBeUndefined();
    });
  });

  // ─── getCell ───────────────────────────────────────────────────────────────

  describe('getCell', () => {
    it('指定日付・タイプのスケジュール配列を返す', () => {
      const s1 = makeSchedule({ id: 's1', order: 1 });
      const s2 = makeSchedule({ id: 's2', order: 2 });
      mockMasterSchedules = [makeDateItem('2024-04-01', 'master', [s1, s2])];

      const { result } = renderHook(() => useScheduleStore());
      const cell = result.current.getCell('2024-04-01', 'master');
      expect(cell).toHaveLength(2);
      expect(cell[0].id).toBe('s1');
    });

    it('存在しないセルは空配列を返す', () => {
      const { result } = renderHook(() => useScheduleStore());
      const cell = result.current.getCell('2099-01-01', 'master');
      expect(cell).toHaveLength(0);
    });
  });

  // ─── applyMove ─────────────────────────────────────────────────────────────

  describe('applyMove - 同タイプ内移動', () => {
    it('setSchedulesByType を呼んでスケジュールを移動する', () => {
      const s = makeSchedule({ id: 's1', startDate: d('2024-04-01'), endDate: d('2024-04-01'), type: 'master' });
      mockMasterSchedules = [makeDateItem('2024-04-01', 'master', [s])];

      const { result } = renderHook(() => useScheduleStore());
      act(() => {
        result.current.applyMove({ schedule: s, newStartDate: d('2024-04-05'), newType: 'master' });
      });

      expect(mockSetSchedulesByType).toHaveBeenCalledWith(
        'master',
        expect.arrayContaining([
          expect.objectContaining({
            date: '2024-04-05',
            type: 'master',
          }),
        ])
      );
    });
  });

  describe('applyMove - タイプ変更を伴う移動', () => {
    it('元タイプと移動先タイプの両方で setSchedulesByType を呼ぶ', () => {
      const s = makeSchedule({ id: 's1', type: 'master' });
      mockMasterSchedules = [makeDateItem('2024-04-01', 'master', [s])];
      mockCustomSchedules = [];

      const { result } = renderHook(() => useScheduleStore());
      act(() => {
        result.current.applyMove({ schedule: s, newStartDate: d('2024-04-05'), newType: 'custom' });
      });

      expect(mockSetSchedulesByType).toHaveBeenCalledWith('master', expect.any(Array));
      expect(mockSetSchedulesByType).toHaveBeenCalledWith('custom', expect.any(Array));
    });
  });

  // ─── applyMoves ────────────────────────────────────────────────────────────

  describe('applyMoves', () => {
    it('複数スケジュールを一括移動し、master/custom 両方を更新する', () => {
      const s1 = makeSchedule({ id: 's1', type: 'master', order: 1 });
      const s2 = makeSchedule({ id: 's2', type: 'custom', order: 1 });
      mockMasterSchedules = [makeDateItem('2024-04-01', 'master', [s1])];
      mockCustomSchedules = [makeDateItem('2024-04-01', 'custom', [s2])];

      const { result } = renderHook(() => useScheduleStore());
      act(() => {
        result.current.applyMoves([
          { schedule: s1, newStartDate: d('2024-04-10'), newType: 'master' },
          { schedule: s2, newStartDate: d('2024-04-10'), newType: 'custom' },
        ]);
      });

      // 必ず master と custom 両方が更新されること
      expect(mockSetSchedulesByType).toHaveBeenCalledWith('master', expect.any(Array));
      expect(mockSetSchedulesByType).toHaveBeenCalledWith('custom', expect.any(Array));
    });
  });

  // ─── reorderCell ───────────────────────────────────────────────────────────

  describe('reorderCell', () => {
    it('セル内スケジュールを並び替え、order を再割当てする', () => {
      const s1 = makeSchedule({ id: 's1', order: 1 });
      const s2 = makeSchedule({ id: 's2', order: 2 });
      const s3 = makeSchedule({ id: 's3', order: 3 });
      mockMasterSchedules = [makeDateItem('2024-04-01', 'master', [s1, s2, s3])];

      const { result } = renderHook(() => useScheduleStore());
      act(() => {
        result.current.reorderCell('2024-04-01', 'master', 0, 2); // s1 を末尾へ
      });

      const call = mockSetSchedulesByType.mock.calls[0];
      const items: Type.ScheduleDateItem[] = call[1];
      const orders = items[0].schedules.map((s) => ({ id: s.id, order: s.order }));
      // s2, s3, s1 の順に order 1,2,3 が割り当てられる
      expect(orders).toEqual([
        { id: 's2', order: 1 },
        { id: 's3', order: 2 },
        { id: 's1', order: 3 },
      ]);
    });

    it('存在しないセルは何もしない', () => {
      const { result } = renderHook(() => useScheduleStore());
      act(() => {
        result.current.reorderCell('2099-01-01', 'master', 0, 1);
      });
      expect(mockSetSchedulesByType).not.toHaveBeenCalled();
    });
  });

  // ─── restoreSnapshot ───────────────────────────────────────────────────────

  describe('restoreSnapshot', () => {
    it('setSchedulesByTypeFunctional を master/custom 両方に対して呼ぶ', () => {
      const snapshot: Type.Schedule[] = [
        makeSchedule({ id: 's1', type: 'master' }),
        makeSchedule({ id: 's2', type: 'custom' }),
      ];

      const { result } = renderHook(() => useScheduleStore());
      act(() => {
        result.current.restoreSnapshot(snapshot);
      });

      expect(mockSetSchedulesByTypeFunctional).toHaveBeenCalledWith('master', expect.any(Function));
      expect(mockSetSchedulesByTypeFunctional).toHaveBeenCalledWith('custom', expect.any(Function));
    });

    it('updater 関数は既存エントリを削除してスナップショットを復元する', () => {
      const existing = makeSchedule({ id: 's1', startDate: d('2024-04-10'), type: 'master' });
      const snapshot = [makeSchedule({ id: 's1', startDate: d('2024-04-01'), type: 'master' })];

      const { result } = renderHook(() => useScheduleStore());
      act(() => {
        result.current.restoreSnapshot(snapshot);
      });

      // master の updater を取り出して実行
      const masterUpdater = mockSetSchedulesByTypeFunctional.mock.calls.find(
        ([type]) => type === 'master'
      )?.[1] as (prev: Type.ScheduleDateItem[]) => Type.ScheduleDateItem[];

      const prev: Type.ScheduleDateItem[] = [makeDateItem('2024-04-10', 'master', [existing])];
      const next = masterUpdater(prev);

      // 復元後は snapshot の日付 (04-01) にあること
      expect(next.find((item) => item.date === '2024-04-01')).toBeDefined();
      // 移動前の日付 (04-10) の s1 は消えていること
      const oldItem = next.find((item) => item.date === '2024-04-10');
      expect(oldItem?.schedules.find((s) => s.id === 's1')).toBeUndefined();
    });

    it('スナップショットが空の場合でも安全に実行できる', () => {
      const { result } = renderHook(() => useScheduleStore());
      expect(() => {
        act(() => {
          result.current.restoreSnapshot([]);
        });
      }).not.toThrow();
    });
  });
});
