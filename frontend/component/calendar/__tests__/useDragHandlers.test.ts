import { describe, it, expect, vi, beforeEach } from 'vitest';
import { renderHook, act } from '@testing-library/react';
import type { DragStartEvent, DragOverEvent, DragEndEvent } from '@dnd-kit/core';

import { useDragHandlers } from '../useDragHandlers';
import type { Type } from '@/type';

// ─── 外部依存をモック ─────────────────────────────────────────────────────────

let mockSelectedIds = new Set<string>();
const mockSetAllSchedules = vi.fn();
vi.mock('@/provider/SelectionContext', () => ({
  useSelection: () => ({
    selectedIds: mockSelectedIds,
    setAllSchedules: mockSetAllSchedules,
  }),
}));

vi.mock('@/provider/ScheduleProvider', () => ({
  useSchedule: () => ({
    masterSchedules: [],
    customSchedules: [],
    setSchedulesByType: vi.fn(),
    setSchedulesByTypeFunctional: vi.fn(),
  }),
}));

// useScheduleStore / useDragState / useSingleDragHandler / useBulkDragHandler をモック
const mockSingleOnDragStart = vi.fn();
const mockSingleOnDragOver = vi.fn();
const mockSingleOnDragEnd = vi.fn().mockResolvedValue(undefined);
vi.mock('../useSingleDragHandler', () => ({
  useSingleDragHandler: () => ({
    onDragStart: mockSingleOnDragStart,
    onDragOver: mockSingleOnDragOver,
    onDragEnd: mockSingleOnDragEnd,
  }),
}));

const mockBulkOnDragStart = vi.fn();
const mockBulkOnDragOver = vi.fn();
const mockBulkOnDragEnd = vi.fn().mockResolvedValue(undefined);
vi.mock('../useBulkDragHandler', () => ({
  useBulkDragHandler: () => ({
    onDragStart: mockBulkOnDragStart,
    onDragOver: mockBulkOnDragOver,
    onDragEnd: mockBulkOnDragEnd,
  }),
}));

let mockPhase = 'idle' as 'idle' | 'single' | 'bulk';
const mockReset = vi.fn();
vi.mock('../useDragState', () => ({
  useDragState: () => ({
    phase: mockPhase,
    activeSchedule: null,
    snapshot: [],
    bulkOrigins: new Map(),
    startSingleDrag: vi.fn(),
    startBulkDrag: vi.fn(),
    reset: mockReset,
  }),
}));

vi.mock('../useScheduleStore', () => ({
  useScheduleStore: () => ({
    findById: vi.fn(),
    getCell: vi.fn(() => []),
    applyMove: vi.fn(),
    applyMoves: vi.fn(),
    reorderCell: vi.fn(),
    restoreSnapshot: vi.fn(),
  }),
}));

// ─── DnD イベント生成ヘルパー ─────────────────────────────────────────────────

const d = (str: string) => new Date(str);

const makeDragStartEvent = (activeId: string): DragStartEvent =>
  ({
    active: { id: activeId, data: { current: undefined }, rect: { current: { initial: null, translated: null } } },
    activatorEvent: new PointerEvent('pointerdown'),
    collisions: null,
    delta: { x: 0, y: 0 },
    over: null,
  }) as unknown as DragStartEvent;

const makeDragOverEvent = (
  activeId: string,
  overId: string,
  date: Date,
  type: Type.ScheduleType
): DragOverEvent =>
  ({
    active: { id: activeId, data: { current: undefined } },
    over: {
      id: overId,
      data: { current: { date, type } },
      rect: { current: { initial: null, translated: null } } as any,
      disabled: false,
    },
    activatorEvent: new PointerEvent('pointerdown'),
    collisions: null,
    delta: { x: 0, y: 0 },
  }) as unknown as DragOverEvent;

const makeDragEndEvent = (
  activeId: string,
  overId: string | null,
  date?: Date,
  type?: Type.ScheduleType
): DragEndEvent =>
  ({
    active: { id: activeId, data: { current: undefined } },
    over: overId
      ? { id: overId, data: { current: { date: date ?? d('2024-04-01'), type: type ?? 'master' } }, disabled: false }
      : null,
    activatorEvent: new PointerEvent('pointerdown'),
    collisions: null,
    delta: { x: 0, y: 0 },
  }) as unknown as DragEndEvent;

// ─── テスト ───────────────────────────────────────────────────────────────────

describe('useDragHandlers', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    mockSelectedIds = new Set();
    mockPhase = 'idle';
  });

  // ─── ルーティング: handleDragStart ──────────────────────────────────────────

  describe('handleDragStart - ルーティング', () => {
    it('選択なし → single.onDragStart へ委譲する', () => {
      const { result } = renderHook(() => useDragHandlers());
      act(() => {
        result.current.handleDragStart(makeDragStartEvent('s1'));
      });
      expect(mockSingleOnDragStart).toHaveBeenCalledWith('s1');
      expect(mockBulkOnDragStart).not.toHaveBeenCalled();
    });

    it('選択が2件以上かつ activeId が含まれる → bulk.onDragStart へ委譲する', () => {
      mockSelectedIds = new Set(['s1', 's2']);
      const { result } = renderHook(() => useDragHandlers());
      act(() => {
        result.current.handleDragStart(makeDragStartEvent('s1'));
      });
      expect(mockBulkOnDragStart).toHaveBeenCalledWith('s1', mockSelectedIds);
      expect(mockSingleOnDragStart).not.toHaveBeenCalled();
    });

    it('選択が1件のみ → bulk でなく single へ委譲する', () => {
      mockSelectedIds = new Set(['s1']);
      const { result } = renderHook(() => useDragHandlers());
      act(() => {
        result.current.handleDragStart(makeDragStartEvent('s1'));
      });
      expect(mockSingleOnDragStart).toHaveBeenCalled();
      expect(mockBulkOnDragStart).not.toHaveBeenCalled();
    });
  });

  // ─── ルーティング: handleDragOver ───────────────────────────────────────────

  describe('handleDragOver - ルーティング', () => {
    it('phase === single → single.onDragOver へ委譲する', () => {
      mockPhase = 'single';
      const { result } = renderHook(() => useDragHandlers());
      act(() => {
        result.current.handleDragOver(makeDragOverEvent('s1', 'cell-master-2024-04-05', d('2024-04-05'), 'master'));
      });
      expect(mockSingleOnDragOver).toHaveBeenCalled();
      expect(mockBulkOnDragOver).not.toHaveBeenCalled();
    });

    it('phase === bulk → bulk.onDragOver へ委譲する', () => {
      mockPhase = 'bulk';
      const { result } = renderHook(() => useDragHandlers());
      act(() => {
        result.current.handleDragOver(makeDragOverEvent('s1', 'cell-master-2024-04-05', d('2024-04-05'), 'master'));
      });
      expect(mockBulkOnDragOver).toHaveBeenCalled();
      expect(mockSingleOnDragOver).not.toHaveBeenCalled();
    });

    it('over が null の場合何もしない', () => {
      mockPhase = 'single';
      const { result } = renderHook(() => useDragHandlers());
      const event = {
        active: { id: 's1', data: { current: undefined } },
        over: null,
      } as unknown as DragOverEvent;
      act(() => {
        result.current.handleDragOver(event);
      });
      expect(mockSingleOnDragOver).not.toHaveBeenCalled();
    });
  });

  // ─── ルーティング: handleDragEnd ────────────────────────────────────────────

  describe('handleDragEnd - ルーティング', () => {
    it('phase === bulk → bulk.onDragEnd へ委譲する', async () => {
      mockPhase = 'bulk';
      const { result } = renderHook(() => useDragHandlers());
      await act(async () => {
        await result.current.handleDragEnd(makeDragEndEvent('s1', 's1'));
      });
      expect(mockBulkOnDragEnd).toHaveBeenCalledWith('s1');
      expect(mockSingleOnDragEnd).not.toHaveBeenCalled();
    });

    it('phase === single → single.onDragEnd へ委譲する', async () => {
      mockPhase = 'single';
      const { result } = renderHook(() => useDragHandlers());
      await act(async () => {
        await result.current.handleDragEnd(makeDragEndEvent('s1', 's1', d('2024-04-01'), 'master'));
      });
      expect(mockSingleOnDragEnd).toHaveBeenCalled();
      expect(mockBulkOnDragEnd).not.toHaveBeenCalled();
    });

    it('dragEnd 後に dragState.reset を呼ぶ', async () => {
      mockPhase = 'single';
      const { result } = renderHook(() => useDragHandlers());
      await act(async () => {
        await result.current.handleDragEnd(makeDragEndEvent('s1', null));
      });
      expect(mockReset).toHaveBeenCalled();
    });
  });
});
