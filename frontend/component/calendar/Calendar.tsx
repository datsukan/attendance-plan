'use client';

import { useMemo, useRef, useEffect, useState } from 'react';

import { Header } from './Header';
import { CalenderDates } from './CalenderDates';
import { SelectionModeBar } from './SelectionModeBar';
import { BulkRemoveConfirmDialog } from '@/component/dialog/remove/BulkRemoveConfirmDialog';

import { usePagePosition } from './usePagePosition';
import { useTargetDate } from './useTargetDate';
import { useDates } from './useDates';
import { useSchedule } from '@/provider/ScheduleProvider';
import { useSelection } from '@/provider/SelectionContext';
import type { Type } from '@/type';

export const Calender = () => {
  const now = useMemo(() => new Date(), []);
  const calendarRef = useRef<HTMLDivElement>(null);

  const { targetDate, targetYear, targetMonth, incrementMonth, decrementMonth, resetMonth } = useTargetDate(now, calendarRef);
  const { weeks, monthCount, addMonthCount } = useDates(targetDate);
  const { initPagePosition, execWhenPageBottom } = usePagePosition();
  const { addSchedule, masterSchedules, customSchedules, removeBulkSchedules } = useSchedule();
  const { selectedIds, clearSelection, isSelectionMode } = useSelection();

  const [isOpenBulkRemoveDialog, setIsOpenBulkRemoveDialog] = useState(false);
  const [bulkRemoveSchedules, setBulkRemoveSchedules] = useState<Type.Schedule[]>([]);

  useEffect(() => {
    initPagePosition();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  useEffect(() => {
    execWhenPageBottom(addMonthCount);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [monthCount]);

  const triggerBulkDelete = () => {
    if (selectedIds.size === 0 || isOpenBulkRemoveDialog) return;
    const allSchedules = [
      ...masterSchedules.flatMap((d) => d.schedules),
      ...customSchedules.flatMap((d) => d.schedules),
    ];
    const targets = allSchedules.filter((s) => selectedIds.has(s.id));
    if (targets.length === 0) return;
    setBulkRemoveSchedules(targets);
    setIsOpenBulkRemoveDialog(true);
  };

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key !== 'Delete' && e.key !== 'Backspace') return;
      e.preventDefault();
      triggerBulkDelete();
    };

    window.addEventListener('keydown', handleKeyDown);
    return () => window.removeEventListener('keydown', handleKeyDown);
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [selectedIds, masterSchedules, customSchedules, isOpenBulkRemoveDialog]);

  const closeBulkRemoveDialog = () => {
    setIsOpenBulkRemoveDialog(false);
    setBulkRemoveSchedules([]);
  };

  const handleBulkRemove = async () => {
    await removeBulkSchedules(bulkRemoveSchedules);
    clearSelection();
  };

  const prev = () => {
    decrementMonth();
    initPagePosition();
  };

  const next = () => {
    incrementMonth();
    initPagePosition();
  };

  const reset = () => {
    resetMonth();
    initPagePosition();
  };

  return (
    <div className="relative" ref={calendarRef}>
      <BulkRemoveConfirmDialog
        schedules={bulkRemoveSchedules}
        isOpen={isOpenBulkRemoveDialog}
        close={closeBulkRemoveDialog}
        remove={handleBulkRemove}
      />
      <div className="sticky top-0 z-10 bg-white">
        <Header year={targetYear} month={targetMonth} prev={prev} next={next} reset={reset} create={addSchedule} />
        {(isSelectionMode || selectedIds.size > 0) && <SelectionModeBar onDelete={triggerBulkDelete} />}
      </div>
      <CalenderDates weeks={weeks} />
    </div>
  );
};
