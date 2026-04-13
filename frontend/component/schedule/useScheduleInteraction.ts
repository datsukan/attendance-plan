import { useDesktopScheduleInteraction } from './useDesktopScheduleInteraction';
import { useMobileScheduleInteraction } from './useMobileScheduleInteraction';

// タッチが主入力デバイスかどうかを判定する（SSR 対応）
const isTouchPrimary = (): boolean =>
  typeof window !== 'undefined' && window.matchMedia('(pointer: coarse)').matches;

export const useScheduleInteraction = (scheduleId: string, onOpen: () => void) => {
  // Rules of Hooks に従い両方を常に呼び出す。返すのはデバイスに応じた一方のみ。
  const desktopHandlers = useDesktopScheduleInteraction(scheduleId, onOpen);
  const mobileHandlers = useMobileScheduleInteraction(scheduleId, onOpen);

  return isTouchPrimary() ? mobileHandlers : desktopHandlers;
};
