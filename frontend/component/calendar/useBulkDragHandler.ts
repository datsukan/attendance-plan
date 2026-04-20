import { useCallback } from 'react';
import { getTime } from 'date-fns';

import { Type } from '@/type';
import { Model } from '@/model';
import { useUndo } from '@/provider/UndoProvider';
import { useSelection } from '@/provider/SelectionContext';
import { updateBulkSchedules } from '@/backend-api/updateBulkSchedules';
import { SessionExpiredError } from '@/backend-api/error';
import { toast } from 'react-hot-toast';

import { buildBulkMoves, hasBulkMoved } from './drag-helpers';
import { useScheduleStore } from './useScheduleStore';
import { useDragState } from './useDragState';

type DragOverPayload = {
  activeId: string;
  targetDate: Date;
  targetType: Type.ScheduleType;
};

/**
 * バルクドラッグの戦略フック。
 *
 * - dragOver：プライマリのみ楽観的に移動（他アイテムは dragEnd でまとめて処理）
 * - dragEnd ：全アイテムを一括移動し、API を1回呼び、Undo を登録
 */
export const useBulkDragHandler = (
  store: ReturnType<typeof useScheduleStore>,
  dragState: ReturnType<typeof useDragState>
) => {
  const { setUndoCommand } = useUndo();
  const { clearSelection } = useSelection();

  // ─── dragStart ─────────────────────────────────────────────────────────────

  const onDragStart = useCallback(
    (activeId: string, selectedIds: Set<string>) => {
      const schedule = store.findById(activeId);
      if (!schedule) return;

      const origins = new Map<string, Date>();
      const fullSnapshot: Type.Schedule[] = [];

      for (const id of selectedIds) {
        const s = store.findById(id);
        if (s) {
          origins.set(id, new Date(s.startDate));
          fullSnapshot.push(s);
        }
      }

      dragState.startBulkDrag(schedule, origins, fullSnapshot);
    },
    [store, dragState]
  );

  // ─── dragOver ──────────────────────────────────────────────────────────────
  // プライマリのみ移動。他アイテムは dragEnd で dateDelta を計算してまとめて処理する。

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
    async (activeId: string) => {
      const { snapshot, bulkOrigins } = dragState;

      const primaryCurrent = store.findById(activeId);
      if (!primaryCurrent) {
        clearSelection();
        return;
      }

      const targetType = primaryCurrent.type;

      // 全選択アイテムの移動後スケジュールを構築
      const movedItems = buildBulkMoves(
        snapshot,
        activeId,
        bulkOrigins,
        targetType,
        primaryCurrent.startDate
      );

      // 実際に動いたかを判定
      const moved = hasBulkMoved(
        activeId,
        bulkOrigins,
        primaryCurrent.startDate,
        targetType,
        snapshot
      );

      // 状態を一括反映
      const moves = movedItems.map((s) => ({
        schedule: store.findById(s.id) ?? s,
        newStartDate: s.startDate,
        newType: s.type,
      }));
      store.applyMoves(moves);

      clearSelection();

      if (!moved) return;

      // API へ一括送信
      try {
        await updateBulkSchedules(movedItems);
      } catch (e) {
        if (e instanceof SessionExpiredError) return;
        toast.error(String(e));
        store.restoreSnapshot(snapshot);
        return;
      }

      // Undo 登録
      setUndoCommand({
        label: `${snapshot.length}件のスケジュールを移動しました`,
        execute: async () => {
          const sorted = [...snapshot].sort((a, b) => a.order - b.order);
          await updateBulkSchedules(sorted);
          store.restoreSnapshot(snapshot);
        },
      });
    },
    [store, dragState, setUndoCommand, clearSelection]
  );

  return { onDragStart, onDragOver, onDragEnd };
};
