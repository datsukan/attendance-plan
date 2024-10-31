import { useDroppable, DragOverlay } from '@dnd-kit/core';
import { arrayMove, SortableContext, sortableKeyboardCoordinates, useSortable } from '@dnd-kit/sortable';

import { ScheduleWeekItem } from '@/component/schedule/ScheduleWeekItem';
import { ScheduleItem } from '@/component/schedule/ScheduleItem';

import { Schedule } from '@/type/schedule';
import { useDateKey } from '@/component/useDateKey';
import {
  isShowItem,
  isDisplaySchedule,
  getColStartClassName,
  getColEndClassName,
  toScheduleTypeName,
} from '@/component/schedule/schedule-module';
import { EditSchedule } from '@/model/edit-schedule';

type Props = {
  type: 'master' | 'custom';
  dates: Date[];
  schedules: Schedule[];
  activeSchedule: Schedule | null;
  removeSchedule: (id: string) => void;
  saveSchedule: (editSchedule: EditSchedule) => void;
  changeScheduleColor: (id: string, color: string) => void;
};

export const ScheduleWeek = ({ type, dates, schedules, activeSchedule, removeSchedule, saveSchedule, changeScheduleColor }: Props) => {
  const { dateToKey } = useDateKey();

  if (!dates || dates.length === 0) {
    return;
  }

  if (!schedules || schedules.length === 0) {
    return;
  }

  return (
    <div className="min-h-14 grid">
      <div className="col-start-1 row-start-1 grid grid-cols-7 h-full">
        {dates.map((date) => (
          <Droppable key={`${type}-${dateToKey(date)}`} id={`${type}-${dateToKey(date)}`} date={date} type={type} />
        ))}
      </div>
      <div className="col-start-1 row-start-1 grid grid-cols-7 gap-y-1 grid-flow-col">
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
    <div ref={setNodeRef} className={isOver ? 'bg-blue-50 flex justify-center items-center' : ''}>
      {isOver && <span className="z-10 text-sm text-blue-600 px-3 py-1 rounded-md bg-blue-50">{toScheduleTypeName(type)}</span>}
    </div>
  );
};
