import { useState } from 'react';

import { Type } from '@/type';

export type DragPhase = 'idle' | 'single' | 'bulk';

type DragState = {
  phase: DragPhase;
  activeSchedule: Type.Schedule | null;
  /** Undo 用スナップショット（単体：移動元セル全体、バルク：選択アイテム全体） */
  snapshot: Type.Schedule[];
  /** バルクドラッグ用：各アイテムのドラッグ開始時の startDate */
  bulkOrigins: Map<string, Date>;
};

const INITIAL_STATE: DragState = {
  phase: 'idle',
  activeSchedule: null,
  snapshot: [],
  bulkOrigins: new Map(),
};

/**
 * ドラッグ中の一時状態を管理するフック。
 *
 * 状態遷移:
 *   idle → (startSingleDrag) → single → (reset) → idle
 *   idle → (startBulkDrag)   → bulk   → (reset) → idle
 */
export const useDragState = () => {
  const [state, setState] = useState<DragState>(INITIAL_STATE);

  const startSingleDrag = (
    activeSchedule: Type.Schedule,
    cellSnapshot: Type.Schedule[]
  ) => {
    setState({
      phase: 'single',
      activeSchedule,
      snapshot: cellSnapshot,
      bulkOrigins: new Map(),
    });
  };

  const startBulkDrag = (
    activeSchedule: Type.Schedule,
    origins: Map<string, Date>,
    fullSnapshot: Type.Schedule[]
  ) => {
    setState({
      phase: 'bulk',
      activeSchedule,
      snapshot: fullSnapshot,
      bulkOrigins: origins,
    });
  };

  const reset = () => setState(INITIAL_STATE);

  return {
    phase: state.phase,
    activeSchedule: state.activeSchedule,
    snapshot: state.snapshot,
    bulkOrigins: state.bulkOrigins,
    startSingleDrag,
    startBulkDrag,
    reset,
  };
};
