import { useState, useEffect, useRef } from 'react';
import { format } from 'date-fns';

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
  const ref = useRef<HTMLDivElement>(null);
  const documentClickHandler = useRef<(this: Document, ev: MouseEvent) => void>();

  const [menuPosition, setMenuPosition] = useState({ x: 0, y: 0 });
  const [infoCardPosition, setInfoCardPosition] = useState({ x: 0, y: 0 });
  const [isOpenMenu, setIsOpenMenu] = useState(false);
  const [isOpenInfoCard, setIsOpenInfoCard] = useState(false);
  const [isOpenRemoveConfirmDialog, setIsOpenRemoveConfirmDialog] = useState(false);
  const [isOpenEditDialog, setIsOpenEditDialog] = useState(false);

  useEffect(() => {
    documentClickHandler.current = (event: MouseEvent) => {
      if (ref.current && !ref.current.contains(event.target as Node)) {
        setIsOpenMenu(false);
        setIsOpenInfoCard(false);

        document.removeEventListener('keydown', handleKeyDown);
        document.removeEventListener('click', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
        document.removeEventListener('contextmenu', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
      }
    };
  }, []);

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

    const rect = event.currentTarget.getBoundingClientRect();
    setMenuPosition({ x: event.clientX - rect.left, y: event.clientY - rect.top });
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

    const rect = event.currentTarget.getBoundingClientRect();
    setInfoCardPosition({ x: event.clientX - rect.left, y: event.clientY - rect.top });
    setIsOpenInfoCard(true);
    setIsOpenMenu(false);

    document.addEventListener('click', documentClickHandler.current as (this: Document, ev: MouseEvent) => any);
    document.addEventListener('keydown', handleKeyDown);
  };

  return (
    <div className="relative">
      <div
        className={`flex h-6 items-center rounded px-2 hover:cursor-pointer ${getColorClassName(schedule.color)}`}
        onContextMenu={onRightClick}
        onClick={onLeftClick}
        ref={ref}
      >
        <span className="truncate text-xs">{generateLabel()}</span>
      </div>
      {isOpenMenu && (
        <div className="absolute z-10 min-w-max" style={{ top: menuPosition.y, left: menuPosition.x }}>
          <Menu
            onSelectColor={(color) => changeScheduleColor(schedule.id, schedule.type, color)}
            openRemoveConfirmDialog={() => setIsOpenRemoveConfirmDialog(true)}
            openEditDialog={() => setIsOpenEditDialog(true)}
          />
        </div>
      )}
      {isOpenInfoCard && (
        <div className="absolute z-10 min-w-max" style={{ top: infoCardPosition.y, left: infoCardPosition.x }}>
          <InfoCard schedule={schedule} />
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
