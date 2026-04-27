import { useEffect, useState } from 'react';
import { useSensor, useSensors, PointerSensor } from '@dnd-kit/core';
import type { DragStartEvent, DragOverEvent, DragEndEvent } from '@dnd-kit/core';
import { getTime } from 'date-fns';

import { Type } from '@/type';
import { useSchedule } from '@/provider/ScheduleProvider';
import { useSelection } from '@/provider/SelectionContext';

import { useScheduleStore } from './useScheduleStore';
import { useDragState } from './useDragState';
import { useSingleDragHandler } from './useSingleDragHandler';
import { useBulkDragHandler } from './useBulkDragHandler';

/**
 * DnD イベントのオーケストレーター。
 *
 * ドラッグの phase（idle / single / bulk）に応じて
 * useSingleDragHandler または useBulkDragHandler へ処理を委譲する。
 *
 * 返り値の sensors / handleDragStart / handleDragOver / handleDragEnd を
 * DndContext へ渡す。
 */
export const useDragHandlers = () => {
  const { masterSchedules, customSchedules } = useSchedule();
  const store = useScheduleStore();
  const dragState = useDragState();
  const { selectedIds, setAllSchedules } = useSelection();
  const [activeDragWidth, setActiveDragWidth] = useState<number | null>(null);

  const single = useSingleDragHandler(store, dragState);
  const bulk = useBulkDragHandler(store, dragState);

  const sensors = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 5 } })
  );

  // 範囲選択・同日チェック用にスケジュール一覧を SelectionContext へ同期
  useEffect(() => {
    const allItems = [...masterSchedules, ...customSchedules];
    allItems.sort((a, b) => {
      if (a.date !== b.date) return a.date.localeCompare(b.date);
      return a.type === 'master' ? -1 : 1;
    });
    const refs = allItems.flatMap((item) =>
      [...item.schedules].sort((a, b) => a.order - b.order).map((s) => ({ id: s.id, date: item.date }))
    );
    setAllSchedules(refs);
  }, [masterSchedules, customSchedules, setAllSchedules]);

  // ─── イベントハンドラ ──────────────────────────────────────────────────────

  const handleDragStart = (event: DragStartEvent) => {
    navigator.vibrate?.(10);

    const activeId = event.active.id.toString();
    const isBulk = selectedIds.has(activeId) && selectedIds.size > 1;

    // ドラッグ開始時点の要素の実際の幅を記録する。
    // DragOverlay の幅をこの値に合わせることで、CSS Grid の列スパンに
    // 対応した正確な幅を再現できる。
    setActiveDragWidth(event.active.rect.current.initial?.width ?? null);

    if (isBulk) {
      bulk.onDragStart(activeId, selectedIds);
    } else {
      single.onDragStart(activeId);
    }
  };

  const handleDragOver = (event: DragOverEvent) => {
    const { active, over } = event;
    if (!over) return;

    const targetDate: Date | undefined = over.data.current?.date;
    const targetType: Type.ScheduleType | undefined = over.data.current?.type;

    if (!targetDate || Number.isNaN(getTime(targetDate)) || !targetType) return;

    const payload = {
      activeId: active.id.toString(),
      overId: over.id.toString(),
      targetDate,
      targetType,
    };

    if (dragState.phase === 'bulk') {
      bulk.onDragOver(payload);
    } else {
      single.onDragOver(payload);
    }
  };

  const handleDragEnd = async (event: DragEndEvent) => {
    const { active, over } = event;

    // ドロップ確定と同時に activeSchedule を null にして opacity-50 を即座に解除する。
    // reset() より先に phase を保存しておくことでルーティングに使える。
    // 戦略ハンドラ側は snapshot / bulkOrigins を最初の await より前にローカル変数へ
    // 捕捉しているため、この時点で reset() を呼んでも非同期処理に影響しない。
    const phase = dragState.phase;
    dragState.reset();
    setActiveDragWidth(null);

    if (phase === 'bulk') {
      await bulk.onDragEnd(active.id.toString());
    } else {
      await single.onDragEnd({
        activeId: active.id.toString(),
        overId: over?.id.toString() ?? null,
      });
    }
  };

  return {
    sensors,
    handleDragStart,
    handleDragOver,
    handleDragEnd,
    activeSchedule: dragState.activeSchedule,
    activeDragWidth,
    bulkCount: dragState.bulkOrigins.size,
  };
};
