import { useState } from 'react';
import { useDroppable, DragOverlay } from '@dnd-kit/core';
import { SortableContext } from '@dnd-kit/sortable';

import { ScheduleWeekItem } from '@/component/schedule/ScheduleWeekItem';
import { ScheduleItem } from '@/component/schedule/ScheduleItem';

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
import { useSchedule } from '@/provider/ScheduleProvider';

type Props = {
  type: Type.ScheduleType;
  dates: Date[];
  schedules: Type.ScheduleDateItem[];
  activeSchedule: Type.Schedule | null;
};

export const ScheduleWeek = ({ type, dates, schedules: scheduleDateItem, activeSchedule }: Props) => {
  const { addSchedule } = useSchedule();
  const [isOpenCreateDialog, setIsOpenCreateDialog] = useState(false);
  const [createDate, setCreateDate] = useState<Date | null>(null);
  const { dateToKey } = useDateKey();
  const schedules = scheduleDateItem.flatMap((item) => item.schedules);

  if (!dates || dates.length === 0) {
    return;
  }

  if (!scheduleDateItem || scheduleDateItem.length === 0) {
    return;
  }

  const openCreateDialog = async (date: Date) => {
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
          return (
            <SortableContext items={displaySchedules} key={`${type}-${dateToKey(date)}`}>
              {displaySchedules.map((schedule) => {
                const colStartClassName = getColStartClassName(index);
                const colEndClassName = getColEndClassName(index, schedule, dates);

                return (
                  <div key={schedule.id} className={`pointer-events-auto mr-2 ${colStartClassName} ${colEndClassName}`}>
                    <ScheduleWeekItem schedule={schedule} isActive={activeSchedule !== null && activeSchedule.id === schedule.id} />
                  </div>
                );
              })}
            </SortableContext>
          );
        })}
      </div>
      {activeSchedule && (
        <DragOverlay>
          <ScheduleItem schedule={activeSchedule} />
        </DragOverlay>
      )}
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
    data: { date, type },
  });

  return (
    <div ref={setNodeRef} className={isOver ? 'flex items-center justify-center bg-blue-200' : ''}>
      {isOver && <span className="z-[9999] rounded-md bg-blue-600 px-3 py-1 text-sm text-blue-50">{toScheduleTypeName(type)}</span>}
    </div>
  );
};
