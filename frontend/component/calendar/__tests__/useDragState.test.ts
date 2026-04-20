import { describe, it, expect } from 'vitest';
import { renderHook, act } from '@testing-library/react';

import { useDragState } from '../useDragState';
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

describe('useDragState', () => {
  it('初期状態は idle', () => {
    const { result } = renderHook(() => useDragState());
    expect(result.current.phase).toBe('idle');
    expect(result.current.activeSchedule).toBeNull();
    expect(result.current.snapshot).toHaveLength(0);
    expect(result.current.bulkOrigins.size).toBe(0);
  });

  describe('startSingleDrag', () => {
    it('phase が single になる', () => {
      const { result } = renderHook(() => useDragState());
      const schedule = makeSchedule();
      act(() => {
        result.current.startSingleDrag(schedule, [schedule]);
      });
      expect(result.current.phase).toBe('single');
    });

    it('activeSchedule と snapshot が設定される', () => {
      const { result } = renderHook(() => useDragState());
      const s1 = makeSchedule({ id: 's1' });
      const s2 = makeSchedule({ id: 's2' });
      act(() => {
        result.current.startSingleDrag(s1, [s1, s2]);
      });
      expect(result.current.activeSchedule?.id).toBe('s1');
      expect(result.current.snapshot).toHaveLength(2);
    });

    it('bulkOrigins は空のまま', () => {
      const { result } = renderHook(() => useDragState());
      act(() => {
        result.current.startSingleDrag(makeSchedule(), []);
      });
      expect(result.current.bulkOrigins.size).toBe(0);
    });
  });

  describe('startBulkDrag', () => {
    it('phase が bulk になる', () => {
      const { result } = renderHook(() => useDragState());
      const origins = new Map([['s1', d('2024-04-01')]]);
      act(() => {
        result.current.startBulkDrag(makeSchedule({ id: 's1' }), origins, []);
      });
      expect(result.current.phase).toBe('bulk');
    });

    it('origins と snapshot が設定される', () => {
      const { result } = renderHook(() => useDragState());
      const origins = new Map([
        ['s1', d('2024-04-01')],
        ['s2', d('2024-04-03')],
      ]);
      const snapshot = [makeSchedule({ id: 's1' }), makeSchedule({ id: 's2' })];
      act(() => {
        result.current.startBulkDrag(makeSchedule({ id: 's1' }), origins, snapshot);
      });
      expect(result.current.bulkOrigins.size).toBe(2);
      expect(result.current.snapshot).toHaveLength(2);
    });
  });

  describe('reset', () => {
    it('single ドラッグ後に idle へ戻る', () => {
      const { result } = renderHook(() => useDragState());
      act(() => {
        result.current.startSingleDrag(makeSchedule(), [makeSchedule()]);
      });
      act(() => {
        result.current.reset();
      });
      expect(result.current.phase).toBe('idle');
      expect(result.current.activeSchedule).toBeNull();
      expect(result.current.snapshot).toHaveLength(0);
    });

    it('bulk ドラッグ後に idle へ戻り、bulkOrigins もクリアされる', () => {
      const { result } = renderHook(() => useDragState());
      const origins = new Map([['s1', d('2024-04-01')]]);
      act(() => {
        result.current.startBulkDrag(makeSchedule(), origins, []);
      });
      act(() => {
        result.current.reset();
      });
      expect(result.current.phase).toBe('idle');
      expect(result.current.bulkOrigins.size).toBe(0);
    });
  });
});
