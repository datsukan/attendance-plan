import { describe, it, expect, vi, beforeEach } from 'vitest';
import { renderHook, act } from '@testing-library/react';

import { useSingleDragHandler } from '../useSingleDragHandler';
import type { Type } from '@/type';

// ─── 外部依存をモック ─────────────────────────────────────────────────────────

const mockSetUndoCommand = vi.fn();
vi.mock('@/provider/UndoProvider', () => ({
  useUndo: () => ({ setUndoCommand: mockSetUndoCommand }),
}));

const mockClearSelection = vi.fn();
vi.mock('@/provider/SelectionContext', () => ({
  useSelection: () => ({ clearSelection: mockClearSelection }),
}));

const mockUpdateBulkSchedules = vi.fn().mockResolvedValue(undefined);
vi.mock('@/backend-api/updateBulkSchedules', () => ({
  updateBulkSchedules: (...args: unknown[]) => mockUpdateBulkSchedules(...args),
}));

vi.mock('@/backend-api/error', () => ({
  SessionExpiredError: class SessionExpiredError extends Error {},
}));

vi.mock('react-hot-toast', () => ({
  toast: { error: vi.fn() },
}));

// ─── store/dragState モック ───────────────────────────────────────────────────

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

const makeStore = (findById: (id: string) => Type.Schedule | undefined = () => undefined) => ({
  findById: vi.fn(findById),
  getCell: vi.fn((_dateKey: string, _type: string) => [] as Type.Schedule[]),
  applyMove: vi.fn(),
  applyMoves: vi.fn(),
  reorderCell: vi.fn(),
  restoreSnapshot: vi.fn(),
});

const makeDragState = (overrides: Partial<ReturnType<typeof makeDragStateBase>> = {}) => ({
  ...makeDragStateBase(),
  ...overrides,
});

function makeDragStateBase() {
  return {
    phase: 'idle' as const,
    activeSchedule: null as Type.Schedule | null,
    snapshot: [] as Type.Schedule[],
    bulkOrigins: new Map<string, Date>(),
    startSingleDrag: vi.fn(),
    startBulkDrag: vi.fn(),
    reset: vi.fn(),
  };
}

// ─── テスト ───────────────────────────────────────────────────────────────────

describe('useSingleDragHandler', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  // ─── onDragStart ────────────────────────────────────────────────────────────

  describe('onDragStart', () => {
    it('スケジュールが存在する場合 startSingleDrag を呼ぶ', () => {
      const schedule = makeSchedule({ id: 's1' });
      const store = makeStore(() => schedule);
      store.getCell.mockReturnValue([schedule]);
      const dragState = makeDragState();

      const { result } = renderHook(() => useSingleDragHandler(store as any, dragState as any));
      act(() => {
        result.current.onDragStart('s1');
      });

      expect(dragState.startSingleDrag).toHaveBeenCalledWith(schedule, [schedule]);
    });

    it('スケジュールが見つからない場合 startSingleDrag を呼ばない', () => {
      const store = makeStore(() => undefined);
      const dragState = makeDragState();

      const { result } = renderHook(() => useSingleDragHandler(store as any, dragState as any));
      act(() => {
        result.current.onDragStart('not-exists');
      });

      expect(dragState.startSingleDrag).not.toHaveBeenCalled();
    });

    it('clearSelection を呼ぶ', () => {
      const schedule = makeSchedule();
      const store = makeStore(() => schedule);
      const dragState = makeDragState();

      const { result } = renderHook(() => useSingleDragHandler(store as any, dragState as any));
      act(() => {
        result.current.onDragStart('schedule-1');
      });

      expect(mockClearSelection).toHaveBeenCalled();
    });
  });

  // ─── onDragOver ─────────────────────────────────────────────────────────────

  describe('onDragOver', () => {
    it('applyMove を呼んで楽観的に状態を更新する', () => {
      const schedule = makeSchedule({ id: 's1' });
      const store = makeStore(() => schedule);
      const dragState = makeDragState();

      const { result } = renderHook(() => useSingleDragHandler(store as any, dragState as any));
      act(() => {
        result.current.onDragOver({
          activeId: 's1',
          targetDate: d('2024-04-05'),
          targetType: 'master',
        });
      });

      expect(store.applyMove).toHaveBeenCalledWith({
        schedule,
        newStartDate: d('2024-04-05'),
        newType: 'master',
      });
    });

    it('スケジュールが見つからない場合 applyMove を呼ばない', () => {
      const store = makeStore(() => undefined);
      const dragState = makeDragState();

      const { result } = renderHook(() => useSingleDragHandler(store as any, dragState as any));
      act(() => {
        result.current.onDragOver({
          activeId: 'not-exists',
          targetDate: d('2024-04-05'),
          targetType: 'master',
        });
      });

      expect(store.applyMove).not.toHaveBeenCalled();
    });
  });

  // ─── onDragEnd ──────────────────────────────────────────────────────────────

  describe('onDragEnd', () => {
    it('overId が null の場合スナップショットを復元する', async () => {
      const schedule = makeSchedule({ id: 's1' });
      const snapshot = [schedule];
      const store = makeStore(() => schedule);
      const dragState = makeDragState({ snapshot, activeSchedule: schedule });

      const { result } = renderHook(() => useSingleDragHandler(store as any, dragState as any));
      await act(async () => {
        await result.current.onDragEnd({
          activeId: 's1',
          overId: null,
          targetDate: null,
          targetType: null,
        });
      });

      expect(store.restoreSnapshot).toHaveBeenCalledWith(snapshot);
      expect(mockUpdateBulkSchedules).not.toHaveBeenCalled();
    });

    it('位置が変わっていない場合は API を呼ばない', async () => {
      const schedule = makeSchedule({ id: 's1', startDate: d('2024-04-01'), type: 'master' });
      const snapshot = [schedule];
      const store = makeStore(() => schedule);
      store.getCell.mockReturnValue([schedule]);
      const dragState = makeDragState({ snapshot, activeSchedule: schedule });

      const { result } = renderHook(() => useSingleDragHandler(store as any, dragState as any));
      await act(async () => {
        await result.current.onDragEnd({
          activeId: 's1',
          overId: 's1', // 同じ ID = 位置変化なし扱い
          targetDate: d('2024-04-01'),
          targetType: 'master',
        });
      });

      expect(mockUpdateBulkSchedules).not.toHaveBeenCalled();
    });

    it('日付が変わった場合は API を呼び Undo を登録する', async () => {
      const originalSchedule = makeSchedule({ id: 's1', startDate: d('2024-04-01'), type: 'master' });
      const movedSchedule = makeSchedule({ id: 's1', startDate: d('2024-04-05'), type: 'master' });
      const snapshot = [originalSchedule];

      let callCount = 0;
      const store = makeStore((id) => {
        // 最初の呼び出しは移動後の状態を返す
        callCount++;
        return callCount === 1 ? movedSchedule : movedSchedule;
      });
      store.getCell.mockReturnValue([movedSchedule]);

      const dragState = makeDragState({ snapshot, activeSchedule: originalSchedule });

      const { result } = renderHook(() => useSingleDragHandler(store as any, dragState as any));
      await act(async () => {
        await result.current.onDragEnd({
          activeId: 's1',
          overId: 's1',
          targetDate: d('2024-04-05'),
          targetType: 'master',
        });
      });

      expect(mockUpdateBulkSchedules).toHaveBeenCalled();
      expect(mockSetUndoCommand).toHaveBeenCalledWith(
        expect.objectContaining({ label: '「テスト」を移動しました' })
      );
    });

    it('同セル内の並び替えで reorderCell を呼ぶ', async () => {
      const s1 = makeSchedule({ id: 's1', startDate: d('2024-04-01'), type: 'master', order: 1 });
      const s2 = makeSchedule({ id: 's2', startDate: d('2024-04-01'), type: 'master', order: 2 });
      const snapshot = [s1, s2];

      const store = makeStore((id) => (id === 's1' ? s1 : s2));
      store.getCell.mockReturnValue([s1, s2]);

      const dragState = makeDragState({ snapshot, activeSchedule: s1 });

      const { result } = renderHook(() => useSingleDragHandler(store as any, dragState as any));
      await act(async () => {
        await result.current.onDragEnd({
          activeId: 's1',
          overId: 's2',
          targetDate: d('2024-04-01'),
          targetType: 'master',
        });
      });

      expect(store.reorderCell).toHaveBeenCalledWith('2024-04-01', 'master', 0, 1);
    });
  });
});
