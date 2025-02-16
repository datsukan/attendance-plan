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
import { Model } from '@/model';

type Props = {
  type: Type.ScheduleType;
  dates: Date[];
  schedules: Type.ScheduleDateItem[];
  activeSchedule: Type.Schedule | null;
  removeSchedule: (id: string, type: Type.ScheduleType) => void;
  saveSchedule: (editSchedule: Model.EditSchedule) => void;
  changeScheduleColor: (id: string, type: Type.ScheduleType, color: string) => void;
};

export const ScheduleWeek = ({
  type,
  dates,
  schedules: scheduleDateItem,
  activeSchedule,
  removeSchedule,
  saveSchedule,
  changeScheduleColor,
}: Props) => {
  const { dateToKey } = useDateKey();
  const schedules = scheduleDateItem.flatMap((item) => item.schedules);

  if (!dates || dates.length === 0) {
    return;
  }

  if (!scheduleDateItem || scheduleDateItem.length === 0) {
    return;
  }

  return (
    <div className="grid min-h-14">
      <div className="col-start-1 row-start-1 grid h-full grid-cols-7">
        {dates.map((date) => (
          <Droppable key={`${type}-${dateToKey(date)}`} id={`${type}-${dateToKey(date)}`} date={date} type={type} />
        ))}
      </div>
      <div className="col-start-1 row-start-1 grid grid-flow-col grid-cols-7 gap-y-1 pb-4">
        {dates.map((date, index) => {
          const displaySchedules = schedules.filter((schedule) => isDisplaySchedule(schedule, date) && isShowItem(index, schedule, date));
          return (
            <SortableContext items={displaySchedules} key={`${type}-${dateToKey(date)}`}>
              {displaySchedules.map((schedule) => {
                const colStartClassName = getColStartClassName(index);
                const colEndClassName = getColEndClassName(index, schedule, dates);

                return (
                  <div key={schedule.id} className={`pr-2 ${colStartClassName} ${colEndClassName}`}>
                    <ScheduleWeekItem
                      schedule={schedule}
                      isActive={activeSchedule !== null && activeSchedule.id === schedule.id}
                      removeSchedule={removeSchedule}
                      saveSchedule={saveSchedule}
                      changeScheduleColor={changeScheduleColor}
                    />
                  </div>
                );
              })}
            </SortableContext>
          );
        })}
      </div>
      {activeSchedule && (
        <DragOverlay>
          <ScheduleItem
            schedule={activeSchedule}
            removeSchedule={removeSchedule}
            saveSchedule={saveSchedule}
            changeScheduleColor={changeScheduleColor}
          />
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
