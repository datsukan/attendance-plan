import { useState, useRef, useMemo } from 'react';
import { format } from 'date-fns';
import { useFloating, useDismiss, useInteractions, autoUpdate, offset, flip, shift, UseFloatingOptions } from '@floating-ui/react';

import { Menu } from './Menu';
import { InfoCard } from './InfoCard';
import { RemoveConfirmDialog } from '@/component/dialog/remove/RemoveConfirmDialog';
import { BulkRemoveConfirmDialog } from '@/component/dialog/remove/BulkRemoveConfirmDialog';
import { EditScheduleDialog } from '@/component/dialog/edit/EditScheduleDialog';

import { hasDateLabel } from '@/component/schedule/schedule-module';
import { getColorClassName } from '@/component/calendar/color-module';
import { useScheduleInteraction } from '@/component/schedule/useScheduleInteraction';
import type { Type } from '@/type';
import { useSchedule } from '@/provider/ScheduleProvider';
import { usePopover } from '@/provider/PopoverProvider';
import { useSelection } from '@/provider/SelectionContext';

type Props = {
  schedule: Type.Schedule;
};

export const ScheduleItem = ({ schedule }: Props) => {
  const { removeSchedule, removeBulkSchedules, masterSchedules, customSchedules, saveSchedule, changeScheduleColor } = useSchedule();
  const { openPopover, closePopover } = usePopover();
  const { selectedIds, toggleSelect, isSelectionMode, clearSelection, enterSelectionMode } = useSelection();
  const documentClickHandler = useRef<(this: Document, ev: MouseEvent) => void>();

  const [isOpenMenu, setIsOpenMenu] = useState(false);
  const [isOpenInfoCard, setIsOpenInfoCard] = useState(false);
  const [isOpenRemoveConfirmDialog, setIsOpenRemoveConfirmDialog] = useState(false);
  const [isOpenBulkRemoveConfirmDialog, setIsOpenBulkRemoveConfirmDialog] = useState(false);
  const [isOpenEditDialog, setIsOpenEditDialog] = useState(false);

  const isBulkTarget = selectedIds.has(schedule.id) && selectedIds.size > 1;

  const allSchedules = useMemo(
    () => [
      ...masterSchedules.flatMap((d) => d.schedules),
      ...customSchedules.flatMap((d) => d.schedules),
    ],
    [masterSchedules, customSchedules]
  );
  const bulkRemoveTargets = useMemo(
    () => (isBulkTarget ? allSchedules.filter((s) => selectedIds.has(s.id)) : []),
    [isBulkTarget, allSchedules, selectedIds]
  );

  const openRemoveDialog = () => {
    if (isBulkTarget) {
      setIsOpenBulkRemoveConfirmDialog(true);
    } else {
      setIsOpenRemoveConfirmDialog(true);
    }
  };

  const floatOptions: UseFloatingOptions = {
    middleware: [offset(10), flip(), shift()],
    whileElementsMounted: autoUpdate,
    placement: 'bottom',
    open: isOpenMenu || isOpenInfoCard,
    onOpenChange: (open) => {
      if (!open) {
        setIsOpenMenu(false);
        setIsOpenInfoCard(false);
        closePopover();
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
      closePopover();
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
      closePopover();
      return;
    }

    setIsOpenMenu(true);
    setIsOpenInfoCard(false);
    openPopover();

    document.addEventListener('click', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
    document.addEventListener('contextmenu', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
    document.addEventListener('keydown', handleKeyDown);
  };

  const openInfoCard = () => {
    clearSelection();

    if (isOpenInfoCard) {
      setIsOpenInfoCard(false);
      setIsOpenMenu(false);
      closePopover();
      return;
    }

    setIsOpenInfoCard(true);
    setIsOpenMenu(false);
    openPopover();

    document.addEventListener('click', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
    document.addEventListener('keydown', handleKeyDown);
  };

  // openInfoCard の定義後に呼び出す（前方参照を避けるため）
  const interactionHandlers = useScheduleInteraction(schedule.id, openInfoCard);

  const onKeyDown = (event: React.KeyboardEvent<HTMLElement>) => {
    if (event.key === 'Enter' || event.key === ' ') {
      event.preventDefault();
      if (isSelectionMode) {
        toggleSelect(schedule.id);
        return;
      }
      openInfoCard();
    }
  };

  return (
    <div className="relative">
      <div
        className={`flex touch-none select-none items-center rounded px-1.5 py-1 hover:cursor-pointer ${getColorClassName(schedule.color)}`}
        onContextMenu={onRightClick}
        role="button"
        ref={refs.setReference}
        {...getReferenceProps({ ...interactionHandlers, onKeyDown, tabIndex: 0 })}
      >
        <span className="line-clamp-2 text-[0.6rem] md:line-clamp-1 md:text-xs">{generateLabel()}</span>
      </div>
      {isOpenMenu && (
        <div ref={refs.setFloating} style={floatingStyles} {...getFloatingProps()} className="z-10 min-w-max">
          <Menu
            onSelectColor={(color) => changeScheduleColor(schedule.id, schedule.type, color)}
            openRemoveConfirmDialog={openRemoveDialog}
            openEditDialog={() => setIsOpenEditDialog(true)}
            onEnterSelectionMode={() => {
              enterSelectionMode(schedule.id);
              setIsOpenMenu(false);
              closePopover();
            }}
          />
        </div>
      )}
      {isOpenInfoCard && (
        <div ref={refs.setFloating} style={floatingStyles} {...getFloatingProps()} className="z-10 min-w-max">
          <InfoCard
            schedule={schedule}
            onSelectColor={(color) => changeScheduleColor(schedule.id, schedule.type, color)}
            openRemoveConfirmDialog={openRemoveDialog}
            openEditDialog={() => setIsOpenEditDialog(true)}
            onEnterSelectionMode={() => {
              enterSelectionMode(schedule.id);
              setIsOpenInfoCard(false);
              closePopover();
            }}
          />
        </div>
      )}
      <RemoveConfirmDialog
        schedule={schedule}
        isOpen={isOpenRemoveConfirmDialog}
        close={() => setIsOpenRemoveConfirmDialog(false)}
        remove={() => removeSchedule(schedule.id, schedule.type)}
      />
      <BulkRemoveConfirmDialog
        schedules={bulkRemoveTargets}
        isOpen={isOpenBulkRemoveConfirmDialog}
        close={() => setIsOpenBulkRemoveConfirmDialog(false)}
        remove={async () => { await removeBulkSchedules(bulkRemoveTargets); clearSelection(); }}
      />
      <EditScheduleDialog schedule={schedule} isOpen={isOpenEditDialog} close={() => setIsOpenEditDialog(false)} submit={saveSchedule} />
    </div>
  );
};
