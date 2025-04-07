import { useState, useRef } from 'react';
import { format } from 'date-fns';
import { useFloating, useDismiss, useInteractions, autoUpdate, offset, flip, shift, UseFloatingOptions } from '@floating-ui/react';

import { Menu } from './Menu';
import { InfoCard } from './InfoCard';
import { RemoveConfirmDialog } from '@/component/dialog/remove/RemoveConfirmDialog';
import { EditScheduleDialog } from '@/component/dialog/edit/EditScheduleDialog';

import { hasDateLabel } from '@/component/schedule/schedule-module';
import { getColorClassName } from '@/component/calendar/color-module';
import type { Type } from '@/type';
import { useSchedule } from '@/provider/ScheduleProvider';

type Props = {
  schedule: Type.Schedule;
};

export const ScheduleItem = ({ schedule }: Props) => {
  const { removeSchedule, saveSchedule, changeScheduleColor } = useSchedule();
  const documentClickHandler = useRef<(this: Document, ev: MouseEvent) => void>();

  const [isOpenMenu, setIsOpenMenu] = useState(false);
  const [isOpenInfoCard, setIsOpenInfoCard] = useState(false);
  const [isOpenRemoveConfirmDialog, setIsOpenRemoveConfirmDialog] = useState(false);
  const [isOpenEditDialog, setIsOpenEditDialog] = useState(false);

  const floatOptions: UseFloatingOptions = {
    middleware: [offset(10), flip(), shift()],
    whileElementsMounted: autoUpdate,
    placement: 'bottom',
    open: isOpenMenu || isOpenInfoCard,
    onOpenChange: (open) => {
      if (!open) {
        setIsOpenMenu(false);
        setIsOpenInfoCard(false);
        document.removeEventListener('keydown', handleKeyDown);
        document.removeEventListener('click', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
        document.removeEventListener('contextmenu', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
      }
    },
  };
  const { refs, floatingStyles, context } = useFloating(floatOptions);
  const dismiss = useDismiss(context);
  const { getReferenceProps, getFloatingProps } = useInteractions([dismiss]);

  const handleKeyDown = (event: KeyboardEvent) => {
    if (event.key === 'Escape') {
      setIsOpenMenu(false);
      setIsOpenInfoCard(false);
    }
  };

  const generateLabel = (): string => {
    const dateFormat = 'M/d';
    if (!hasDateLabel(schedule)) {
      return schedule.name;
    }

    return `${schedule.name} (${format(schedule.startDate, dateFormat)} ~ ${format(schedule.endDate, dateFormat)})`;
  };

  const onRightClick = (event: React.MouseEvent<HTMLElement>) => {
    event.preventDefault();

    if (isOpenMenu) {
      setIsOpenMenu(false);
      setIsOpenInfoCard(false);
      return;
    }

    setIsOpenMenu(true);
    setIsOpenInfoCard(false);

    document.addEventListener('click', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
    document.addEventListener('contextmenu', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
    document.addEventListener('keydown', handleKeyDown);
  };

  const onLeftClick = (event: React.MouseEvent<HTMLElement>) => {
    event.preventDefault();

    if (isOpenInfoCard) {
      setIsOpenInfoCard(false);
      setIsOpenMenu(false);
      return;
    }

    setIsOpenInfoCard(true);
    setIsOpenMenu(false);

    document.addEventListener('click', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
    document.addEventListener('keydown', handleKeyDown);
  };

  return (
    <div className="relative">
      <div
        className={`flex touch-none items-center rounded px-1.5 py-1 hover:cursor-pointer ${getColorClassName(schedule.color)}`}
        onContextMenu={onRightClick}
        onClick={onLeftClick}
        ref={refs.setReference}
        {...getReferenceProps()}
      >
        <span className="line-clamp-2 text-[0.6rem] md:line-clamp-1 md:text-xs">{generateLabel()}</span>
      </div>
      {isOpenMenu && (
        <div ref={refs.setFloating} style={floatingStyles} {...getFloatingProps()} className="z-10 min-w-max">
          <Menu
            onSelectColor={(color) => changeScheduleColor(schedule.id, schedule.type, color)}
            openRemoveConfirmDialog={() => setIsOpenRemoveConfirmDialog(true)}
            openEditDialog={() => setIsOpenEditDialog(true)}
          />
        </div>
      )}
      {isOpenInfoCard && (
        <div ref={refs.setFloating} style={floatingStyles} {...getFloatingProps()} className="z-10 min-w-max">
          <InfoCard
            schedule={schedule}
            onSelectColor={(color) => changeScheduleColor(schedule.id, schedule.type, color)}
            openRemoveConfirmDialog={() => setIsOpenRemoveConfirmDialog(true)}
            openEditDialog={() => setIsOpenEditDialog(true)}
          />
        </div>
      )}
      <RemoveConfirmDialog
        schedule={schedule}
        isOpen={isOpenRemoveConfirmDialog}
        close={() => setIsOpenRemoveConfirmDialog(false)}
        remove={() => removeSchedule(schedule.id, schedule.type)}
      />
      <EditScheduleDialog schedule={schedule} isOpen={isOpenEditDialog} close={() => setIsOpenEditDialog(false)} submit={saveSchedule} />
    </div>
  );
};
