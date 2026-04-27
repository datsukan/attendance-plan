import { useRef } from 'react';
import { useSelection } from '@/provider/SelectionContext';

// dnd-kit の PointerSensor と同じ距離でキャンセルし、ドラッグと競合しないようにする
const LONG_PRESS_DURATION = 500;
const MOVE_THRESHOLD_PX = 5;

type Handlers = {
  onClick: (e: React.MouseEvent<HTMLElement>) => void;
  onContextMenu: (e: React.MouseEvent<HTMLElement>) => void;
  onPointerDown: (e: React.PointerEvent<HTMLElement>) => void;
  onPointerMove: (e: React.PointerEvent<HTMLElement>) => void;
  onPointerUp: (e: React.PointerEvent<HTMLElement>) => void;
  onPointerCancel: (e: React.PointerEvent<HTMLElement>) => void;
};

export const useMobileScheduleInteraction = (scheduleId: string, onOpen: () => void): Handlers => {
  const { toggleSelect, isSelectionMode, enterSelectionMode } = useSelection();

  const timerRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  const startPosRef = useRef<{ x: number; y: number } | null>(null);
  const didLongPressRef = useRef(false);
  const didDragRef = useRef(false);

  const cancelTimer = () => {
    if (timerRef.current) {
      clearTimeout(timerRef.current);
      timerRef.current = null;
    }
  };

  const onPointerDown = (e: React.PointerEvent<HTMLElement>) => {
    // マウス右クリックは対象外
    if (e.pointerType === 'mouse' && e.button !== 0) return;

    startPosRef.current = { x: e.clientX, y: e.clientY };
    didLongPressRef.current = false;
    didDragRef.current = false;

    timerRef.current = setTimeout(() => {
      timerRef.current = null;
      didLongPressRef.current = true;
      navigator.vibrate?.(30);

      if (isSelectionMode) {
        toggleSelect(scheduleId);
      } else {
        enterSelectionMode(scheduleId);
      }
    }, LONG_PRESS_DURATION);
  };

  const onPointerMove = (e: React.PointerEvent<HTMLElement>) => {
    if (!startPosRef.current || !timerRef.current) return;
    const dx = e.clientX - startPosRef.current.x;
    const dy = e.clientY - startPosRef.current.y;
    if (dx * dx + dy * dy >= MOVE_THRESHOLD_PX * MOVE_THRESHOLD_PX) {
      cancelTimer();
      didDragRef.current = true;
    }
  };

  const onPointerUp = () => {
    cancelTimer();
    startPosRef.current = null;
  };

  // dnd-kit がポインターをキャプチャした際などにも確実にタイマーをキャンセルする
  const onPointerCancel = () => {
    cancelTimer();
    startPosRef.current = null;
  };

  // 長押しまたはドラッグ後に発火する contextmenu を抑止してメニューが開かないようにする
  const onContextMenu = (e: React.MouseEvent<HTMLElement>) => {
    if (didLongPressRef.current || didDragRef.current) {
      e.preventDefault();
    }
  };

  const onClick = (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault();

    // 長押しで発火済みの場合は click イベントの処理をスキップ
    if (didLongPressRef.current) {
      didLongPressRef.current = false;
      return;
    }

    // ドラッグ後にモバイルブラウザが発火する spurious click を無視する
    if (didDragRef.current) {
      didDragRef.current = false;
      return;
    }

    if (isSelectionMode) {
      toggleSelect(scheduleId);
      return;
    }

    onOpen();
  };

  return { onClick, onContextMenu, onPointerDown, onPointerMove, onPointerUp, onPointerCancel };
};
