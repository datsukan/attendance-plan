import { useState } from 'react';
import { isBefore, isEqual } from 'date-fns';
import { useDroppable } from '@dnd-kit/core';
import { SortableContext, type SortingStrategy } from '@dnd-kit/sortable';

// CSS Grid のカラムスパンが異なるアイテム同士を CSS transform で入れ替えると
// 幅の表示が崩れる。同セル内の並び替えは onDragOver でのライブ状態更新で処理するため、
// SortableContext の transform ベースの変位アニメーションを無効化する。
const noopSortingStrategy: SortingStrategy = () => null;

import { ScheduleItem } from '@/component/schedule/ScheduleItem';

import { ScheduleWeekItem } from '@/component/schedule/ScheduleWeekItem';

import { Type } from '@/type';
import { useDateKey } from '@/component/useDateKey';
import {
  isShowItem,
  isDisplaySchedule,
  getColStartClassName,
  getColEndClassName,
  toScheduleTypeName,
} from '@/component/schedule/schedule-module';
import { CreateScheduleDialog } from '@/component/dialog/create/CreateScheduleDialog';
import { isDialogRecentlyClosed, isPopoverRecentlyClosed } from '@/component/dialog/close-guard';
import { useSchedule } from '@/provider/ScheduleProvider';
import { usePopover } from '@/provider/PopoverProvider';
import { useSelection } from '@/provider/SelectionContext';

type Props = {
  type: Type.ScheduleType;
  dates: Date[];
  schedules: Type.ScheduleDateItem[];
  activeSchedule: Type.Schedule | null;
};

export const ScheduleWeek = ({ type, dates, schedules: scheduleDateItem, activeSchedule }: Props) => {
  const { addSchedule } = useSchedule();
  const { isOpenPopover } = usePopover();
  const { clearSelection } = useSelection();
  const [isOpenCreateDialog, setIsOpenCreateDialog] = useState(false);
  const [createDate, setCreateDate] = useState<Date | null>(null);
  const { dateToKey } = useDateKey();
  const schedules = scheduleDateItem.flatMap((item) => item.schedules).sort((a, b) => {
    if (a.startDate < b.startDate) return -1;
    if (a.startDate > b.startDate) return 1;
    return a.order - b.order;
  });

  if (!dates || dates.length === 0) {
    return;
  }

  const openCreateDialog = async (date: Date) => {
    clearSelection();
    if (isOpenPopover || isDialogRecentlyClosed() || isPopoverRecentlyClosed()) {
      return;
    }

    setCreateDate(date);
    setIsOpenCreateDialog(true);
  };

  const closeCreateDialog = () => {
    setCreateDate(null);
    setIsOpenCreateDialog(false);
  };

  return (
    <div className="grid min-h-14">
      <CreateScheduleDialog
        defaultType={type}
        defaultDate={createDate}
        isOpen={isOpenCreateDialog}
        close={closeCreateDialog}
        submit={addSchedule}
      />
      <div className="col-start-1 row-start-1 grid h-full grid-cols-7">
        {dates.map((date) => (
          <Droppable key={`${type}-${dateToKey(date)}`} id={`${type}-${dateToKey(date)}`} date={date} type={type} />
        ))}
      </div>
      <div className="col-start-1 row-start-1 grid h-full grid-cols-7">
        {dates.map((date) => {
          return <div key={dateToKey(date)} onClick={() => openCreateDialog(date)} />;
        })}
      </div>
      <div className="pointer-events-none col-start-1 row-start-1 mb-6 grid h-fit grid-flow-col grid-cols-7 gap-y-1">
        {dates.map((date, index) => {
          const displaySchedules = schedules.filter((schedule) => isDisplaySchedule(schedule, date) && isShowItem(index, schedule, date));
          // SortableContext には「この日が startDate のスケジュール」のみ渡す。
          // 週跨ぎ継続表示（startDate < date）を含めると同一 ID が複数の
          // SortableContext に登録されて dnd-kit が壊れる。
          const sortableIds = displaySchedules.filter((s) => isEqual(s.startDate, date)).map((s) => s.id);
          return (
            <SortableContext items={sortableIds} key={`${type}-${dateToKey(date)}`} strategy={noopSortingStrategy}>
              {displaySchedules.map((schedule) => {
                const colStartClassName = getColStartClassName(index);
                const colEndClassName = getColEndClassName(schedule, dates);
                const isContinuation = isBefore(schedule.startDate, date);
                return isContinuation ? (
                  // 週跨ぎ継続: 視覚表示は維持するが useSortable 二重登録を防ぐ
                  <div key={schedule.id} className={`pointer-events-auto mr-2 ${colStartClassName} ${colEndClassName}`}>
                    <ScheduleItem schedule={schedule} />
                  </div>
                ) : (
                  <ScheduleWeekItem
                    key={schedule.id}
                    schedule={schedule}
                    isActive={activeSchedule !== null && activeSchedule.id === schedule.id}
                    colStartClassName={colStartClassName}
                    colEndClassName={colEndClassName}
                  />
                );
              })}
            </SortableContext>
          );
        })}
      </div>
    </div>
  );
};

type DroppableProps = {
  id: string;
  date: Date;
  type: 'master' | 'custom';
};

const Droppable = ({ id, date, type }: DroppableProps) => {
  const { isOver, setNodeRef } = useDroppable({
    id: id,
    data: { date, type, isDroppableBackground: true },
  });

  return (
    <div ref={setNodeRef} className={isOver ? 'flex items-center justify-center bg-blue-200' : ''}>
      {isOver && <span className="z-[9999] rounded-md bg-blue-600 px-3 py-1 text-sm text-blue-50">{toScheduleTypeName(type)}</span>}
    </div>
  );
};
