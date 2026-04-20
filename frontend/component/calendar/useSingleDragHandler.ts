import { useCallback } from 'react';
import { getTime } from 'date-fns';

import { Type } from '@/type';
import { Model } from '@/model';
import { useUndo } from '@/provider/UndoProvider';
import { useSelection } from '@/provider/SelectionContext';
import { updateBulkSchedules } from '@/backend-api/updateBulkSchedules';
import { SessionExpiredError } from '@/backend-api/error';
import { toast } from 'react-hot-toast';

import { hasScheduleMoved } from './drag-helpers';
import { useScheduleStore } from './useScheduleStore';
import { useDragState } from './useDragState';

type DragOverPayload = {
  activeId: string;
  targetDate: Date;
  targetType: Type.ScheduleType;
};

type DragEndPayload = {
  activeId: string;
  overId: string | null;
  targetDate: Date | null;
  targetType: Type.ScheduleType | null;
};

/**
 * 単体ドラッグの戦略フック。
 *
 * - dragOver：楽観的 UI 更新のみ（API 呼び出しなし）
 * - dragEnd ：位置変化を確認したうえで API を1回呼び、Undo を登録
 */
export const useSingleDragHandler = (
  store: ReturnType<typeof useScheduleStore>,
  dragState: ReturnType<typeof useDragState>
) => {
  const { setUndoCommand } = useUndo();
  const { clearSelection } = useSelection();

  // ─── dragStart ─────────────────────────────────────────────────────────────

  const onDragStart = useCallback(
    (activeId: string) => {
      const schedule = store.findById(activeId);
      if (!schedule) return;

      clearSelection();
      const dateKey = Model.ScheduleDateItem.toKey(schedule.startDate);
      const cellSnapshot = store.getCell(dateKey, schedule.type);
      dragState.startSingleDrag(schedule, cellSnapshot);
    },
    [store, dragState, clearSelection]
  );

  // ─── dragOver ──────────────────────────────────────────────────────────────

  const onDragOver = useCallback(
    ({ activeId, targetDate, targetType }: DragOverPayload) => {
      if (!targetDate || Number.isNaN(getTime(targetDate))) return;

      const schedule = store.findById(activeId);
      if (!schedule) return;

      store.applyMove({ schedule, newStartDate: targetDate, newType: targetType });
    },
    [store]
  );

  // ─── dragEnd ───────────────────────────────────────────────────────────────

  const onDragEnd = useCallback(
    async ({ activeId, overId, targetDate, targetType }: DragEndPayload) => {
      const { snapshot, activeSchedule } = dragState;

      // drop 先なし → キャンセル扱い（UI はスナップショットで復元）
      if (!overId) {
        if (snapshot.length > 0) {
          store.restoreSnapshot(snapshot);
        }
        return;
      }

      const current = store.findById(activeId);
      const snapshotItem = snapshot.find((s) => s.id === activeId);

      if (!current || !snapshotItem) return;

      // 同セル内での並び替え
      if (activeId !== overId && targetDate && targetType) {
        const dateKey = Model.ScheduleDateItem.toKey(current.startDate);
        const cell = store.getCell(dateKey, current.type);
        const fromIndex = cell.findIndex((s) => s.id === activeId);
        const toIndex = cell.findIndex((s) => s.id === overId);

        if (fromIndex !== -1 && toIndex !== -1) {
          store.reorderCell(dateKey, current.type, fromIndex, toIndex);

          // 並び替え後の最新セルを API へ
          const reorderedCell = store.getCell(dateKey, current.type);
          await persistAndUndo(
            `「${snapshotItem.name}」を移動しました`,
            snapshot,
            reorderedCell
          );
          return;
        }
      }

      // 日付・タイプの変更ドラッグ（handleDragOver で移動済み）
      if (hasScheduleMoved(snapshotItem, current)) {
        const dateKey = Model.ScheduleDateItem.toKey(current.startDate);
        const cell = store.getCell(dateKey, current.type);
        await persistAndUndo(
          `「${snapshotItem.name}」を移動しました`,
          snapshot,
          cell
        );
      }
    },
    [store, dragState, setUndoCommand]
  );

  // ─── 共通：API呼び出し + Undo登録 ─────────────────────────────────────────

  const persistAndUndo = useCallback(
    async (label: string, snapshot: Type.Schedule[], cellSchedules: Type.Schedule[]) => {
      try {
        await updateBulkSchedules(cellSchedules);
      } catch (e) {
        if (e instanceof SessionExpiredError) return;
        toast.error(String(e));
        store.restoreSnapshot(snapshot);
        return;
      }

      setUndoCommand({
        label,
        execute: async () => {
          const sorted = [...snapshot].sort((a, b) => a.order - b.order);
          await updateBulkSchedules(sorted);
          store.restoreSnapshot(snapshot);
        },
      });
    },
    [store, setUndoCommand]
  );

  return { onDragStart, onDragOver, onDragEnd };
};
