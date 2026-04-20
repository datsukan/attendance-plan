import { describe, it, expect, vi, beforeEach } from 'vitest';
import { renderHook, act } from '@testing-library/react';

import { useBulkDragHandler } from '../useBulkDragHandler';
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

const makeStore = (findById: (id: string) => Type.Schedule | undefined = () => undefined) => ({
  findById: vi.fn(findById),
  getCell: vi.fn(() => [] as Type.Schedule[]),
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

describe('useBulkDragHandler', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  // ─── onDragStart ────────────────────────────────────────────────────────────

  describe('onDragStart', () => {
    it('選択済みアイテムの origins と snapshot を構築して startBulkDrag を呼ぶ', () => {
      const s1 = makeSchedule({ id: 's1', startDate: d('2024-04-01') });
      const s2 = makeSchedule({ id: 's2', startDate: d('2024-04-03') });
      const store = makeStore((id) => (id === 's1' ? s1 : id === 's2' ? s2 : undefined));
      const dragState = makeDragState();

      const { result } = renderHook(() => useBulkDragHandler(store as any, dragState as any));
      act(() => {
        result.current.onDragStart('s1', new Set(['s1', 's2']));
      });

      expect(dragState.startBulkDrag).toHaveBeenCalledWith(
        s1,
        expect.any(Map),
        expect.arrayContaining([
          expect.objectContaining({ id: 's1' }),
          expect.objectContaining({ id: 's2' }),
        ])
      );

      const callArgs = (dragState.startBulkDrag as ReturnType<typeof vi.fn>).mock.calls[0];
      const origins: Map<string, Date> = callArgs[1];
      expect(origins.get('s1')).toBeDefined();
      expect(origins.get('s2')).toBeDefined();
    });

    it('activeId のスケジュールが存在しない場合 startBulkDrag を呼ばない', () => {
      const store = makeStore(() => undefined);
      const dragState = makeDragState();

      const { result } = renderHook(() => useBulkDragHandler(store as any, dragState as any));
      act(() => {
        result.current.onDragStart('not-exists', new Set(['not-exists']));
      });

      expect(dragState.startBulkDrag).not.toHaveBeenCalled();
    });
  });

  // ─── onDragOver ─────────────────────────────────────────────────────────────

  describe('onDragOver', () => {
    it('プライマリのみ applyMove する', () => {
      const schedule = makeSchedule({ id: 's1' });
      const store = makeStore(() => schedule);
      const dragState = makeDragState();

      const { result } = renderHook(() => useBulkDragHandler(store as any, dragState as any));
      act(() => {
        result.current.onDragOver({
          activeId: 's1',
          targetDate: d('2024-04-10'),
          targetType: 'custom',
        });
      });

      expect(store.applyMove).toHaveBeenCalledWith({
        schedule,
        newStartDate: d('2024-04-10'),
        newType: 'custom',
      });
    });
  });

  // ─── onDragEnd ──────────────────────────────────────────────────────────────

  describe('onDragEnd', () => {
    it('実際に移動があった場合、API を呼び Undo を登録する', async () => {
      const s1 = makeSchedule({ id: 's1', startDate: d('2024-04-05'), type: 'master' });
      const s2 = makeSchedule({ id: 's2', startDate: d('2024-04-07'), type: 'master' });

      // snapshot: 移動前の状態
      const snapshot = [
        makeSchedule({ id: 's1', startDate: d('2024-04-01'), type: 'master' }),
        makeSchedule({ id: 's2', startDate: d('2024-04-03'), type: 'master' }),
      ];

      // bulkOrigins: ドラッグ開始時の startDate
      const bulkOrigins = new Map([
        ['s1', d('2024-04-01')],
        ['s2', d('2024-04-03')],
      ]);

      const store = makeStore((id) => (id === 's1' ? s1 : id === 's2' ? s2 : undefined));
      const dragState = makeDragState({ snapshot, bulkOrigins });

      const { result } = renderHook(() => useBulkDragHandler(store as any, dragState as any));
      await act(async () => {
        await result.current.onDragEnd('s1');
      });

      expect(store.applyMoves).toHaveBeenCalled();
      expect(mockUpdateBulkSchedules).toHaveBeenCalled();
      expect(mockSetUndoCommand).toHaveBeenCalledWith(
        expect.objectContaining({ label: '2件のスケジュールを移動しました' })
      );
      expect(mockClearSelection).toHaveBeenCalled();
    });

    it('移動がなかった場合（同じ位置に戻した）は API を呼ばない', async () => {
      const s1 = makeSchedule({ id: 's1', startDate: d('2024-04-01'), type: 'master' });
      const snapshot = [makeSchedule({ id: 's1', startDate: d('2024-04-01'), type: 'master' })];
      const bulkOrigins = new Map([['s1', d('2024-04-01')]]);

      const store = makeStore(() => s1);
      const dragState = makeDragState({ snapshot, bulkOrigins });

      const { result } = renderHook(() => useBulkDragHandler(store as any, dragState as any));
      await act(async () => {
        await result.current.onDragEnd('s1');
      });

      expect(mockUpdateBulkSchedules).not.toHaveBeenCalled();
      expect(mockSetUndoCommand).not.toHaveBeenCalled();
      expect(mockClearSelection).toHaveBeenCalled();
    });

    it('API エラー時はスナップショットを復元し Undo を登録しない', async () => {
      const s1 = makeSchedule({ id: 's1', startDate: d('2024-04-05'), type: 'master' });
      const snapshot = [makeSchedule({ id: 's1', startDate: d('2024-04-01'), type: 'master' })];
      const bulkOrigins = new Map([['s1', d('2024-04-01')]]);

      mockUpdateBulkSchedules.mockRejectedValueOnce(new Error('Network Error'));

      const store = makeStore(() => s1);
      const dragState = makeDragState({ snapshot, bulkOrigins });

      const { result } = renderHook(() => useBulkDragHandler(store as any, dragState as any));
      await act(async () => {
        await result.current.onDragEnd('s1');
      });

      expect(store.restoreSnapshot).toHaveBeenCalledWith(snapshot);
      expect(mockSetUndoCommand).not.toHaveBeenCalled();
    });

    it('primaryId のスケジュールが見つからない場合 clearSelection のみ呼ぶ', async () => {
      const store = makeStore(() => undefined);
      const dragState = makeDragState();

      const { result } = renderHook(() => useBulkDragHandler(store as any, dragState as any));
      await act(async () => {
        await result.current.onDragEnd('not-exists');
      });

      expect(mockClearSelection).toHaveBeenCalled();
      expect(store.applyMoves).not.toHaveBeenCalled();
    });
  });
});
