import { useCallback, useRef } from 'react';
import { getTime, isSameDay } from 'date-fns';

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
  overId?: string;
  targetDate: Date;
  targetType: Type.ScheduleType;
};

type DragEndPayload = {
  activeId: string;
  overId: string | null;
};

/**
 * 単体ドラッグの戦略フック。
 *
 * - dragOver：楽観的 UI 更新のみ（API 呼び出しなし）
 *   - 同セル内：hoverId が変わるたびに store.reorderCell() でライブ並び替え
 *   - セル跨ぎ：store.applyMove() で日付・タイプを楽観的に変更
 * - dragEnd ：位置変化を確認したうえで API を1回呼び、Undo を登録
 */
export const useSingleDragHandler = (
  store: ReturnType<typeof useScheduleStore>,
  dragState: ReturnType<typeof useDragState>
) => {
  const { setUndoCommand } = useUndo();
  const { clearSelection } = useSelection();

  // 同セル内ライブ並び替えの状態管理（ref なのでレンダーを引き起こさない）
  const lastSameCellOverRef = useRef<string | null>(null);
  const sameCellReorderedRef = useRef(false);

  // ─── dragStart ─────────────────────────────────────────────────────────────

  const onDragStart = useCallback(
    (activeId: string) => {
      lastSameCellOverRef.current = null;
      sameCellReorderedRef.current = false;

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
    ({ activeId, overId, targetDate, targetType }: DragOverPayload) => {
      if (!targetDate || Number.isNaN(getTime(targetDate))) return;

      const schedule = store.findById(activeId);
      if (!schedule) return;

      if (isSameDay(targetDate, schedule.startDate) && targetType === schedule.type) {
        // 同セル内：別スケジュール上にポインターが入ったときだけライブ並び替え。
        // lastSameCellOverRef で連続イベントの重複実行を防ぐ。
        if (overId && overId !== activeId && overId !== lastSameCellOverRef.current) {
          const overSchedule = store.findById(overId);
          if (overSchedule) {
            const dateKey = Model.ScheduleDateItem.toKey(schedule.startDate);
            const cell = store.getCell(dateKey, schedule.type);
            const fromIndex = cell.findIndex((s) => s.id === activeId);
            const toIndex = cell.findIndex((s) => s.id === overId);
            if (fromIndex !== -1 && toIndex !== -1 && fromIndex !== toIndex) {
              store.reorderCell(dateKey, schedule.type, fromIndex, toIndex);
              sameCellReorderedRef.current = true;
            }
          }
          lastSameCellOverRef.current = overId;
        }
        return;
      }

      // セル跨ぎ移動：同セル追跡をリセットして楽観的移動を適用
      lastSameCellOverRef.current = null;
      store.applyMove({ schedule, newStartDate: targetDate, newType: targetType });
    },
    [store]
  );

  // ─── dragEnd ───────────────────────────────────────────────────────────────

  const onDragEnd = useCallback(
    async ({ activeId, overId }: DragEndPayload) => {
      const { snapshot } = dragState;
      const didSameCellReorder = sameCellReorderedRef.current;
      sameCellReorderedRef.current = false;
      lastSameCellOverRef.current = null;

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

      // onDragOver でライブ並び替え済み → 現在のセル状態をそのまま永続化
      if (didSameCellReorder) {
        const dateKey = Model.ScheduleDateItem.toKey(current.startDate);
        const cell = store.getCell(dateKey, current.type);
        await persistAndUndo(
          `「${snapshotItem.name}」を移動しました`,
          snapshot,
          cell
        );
        return;
      }

      // 同セル内フォールバック：onDragOver でライブ並び替えが行われなかった場合
      if (activeId !== overId) {
        const dateKey = Model.ScheduleDateItem.toKey(current.startDate);
        const cell = store.getCell(dateKey, current.type);
        const fromIndex = cell.findIndex((s) => s.id === activeId);
        const toIndex = cell.findIndex((s) => s.id === overId);

        if (fromIndex !== -1 && toIndex !== -1) {
          const reorderedCell = store.reorderCell(dateKey, current.type, fromIndex, toIndex);
          if (!reorderedCell) return;
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
