import { useSelection } from '@/provider/SelectionContext';

type Handlers = {
  onClick: (e: React.MouseEvent<HTMLElement>) => void;
};

export const useDesktopScheduleInteraction = (scheduleId: string, onOpen: () => void): Handlers => {
  const { toggleSelect, rangeSelect } = useSelection();

  const onClick = (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault();

    if (e.ctrlKey || e.metaKey) {
      toggleSelect(scheduleId);
      return;
    }

    if (e.shiftKey) {
      rangeSelect(scheduleId);
      return;
    }

    onOpen();
  };

  return { onClick };
};
