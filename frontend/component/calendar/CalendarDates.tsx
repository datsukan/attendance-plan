import { DndContext, DragOverlay, pointerWithin } from '@dnd-kit/core';
import { format } from 'date-fns';

import { CalendarDateItem } from './CalendarDateItem';
import { ScheduleWeekContainer } from '@/component/schedule/ScheduleWeekContainer';
import { ScheduleItem } from '@/component/schedule/ScheduleItem';

import { useDateKey } from '@/component/useDateKey';
import { getColorClassName } from '@/component/calendar/color-module';
import { hasDateLabel } from '@/component/schedule/schedule-module';
import { useDragHandlers } from './useDragHandlers';

type Props = {
  weeks: Date[][];
};

export const CalendarDates = ({ weeks }: Props) => {
  const { dateToKey } = useDateKey();
  const {
    sensors,
    handleDragStart,
    handleDragOver,
    handleDragEnd,
    activeSchedule,
    bulkCount,
  } = useDragHandlers();

  const generateDragLabel = (): string => {
    if (!activeSchedule) return '';
    const dateFormat = 'M/d';
    if (!hasDateLabel(activeSchedule)) {
      return activeSchedule.name;
    }
    return `${activeSchedule.name} (${format(activeSchedule.startDate, dateFormat)} ~ ${format(activeSchedule.endDate, dateFormat)})`;
  };

  return (
    <DndContext
      onDragStart={handleDragStart}
      onDragEnd={handleDragEnd}
      onDragOver={handleDragOver}
      collisionDetection={pointerWithin}
      sensors={sensors}
    >
      <div className="border-b border-r">
        {weeks.map((week) => {
          return (
            <div key={dateToKey(week[0]) + '-frame-week'} className="grid">
              <div data-date={dateToKey(week[0])} className=" calendar-item col-start-1 row-start-1 grid grid-cols-7">
                {week.map((date) => {
                  return (
                    <div key={dateToKey(date) + '-frame-date'} className="border-l border-t">
                      <CalendarDateItem date={date} />
                    </div>
                  );
                })}
              </div>
              <div className="col-start-1 row-start-1 mt-10 pb-1">
                <ScheduleWeekContainer week={week} activeSchedule={activeSchedule} />
              </div>
            </div>
          );
        })}
      </div>
      {activeSchedule && (
        <DragOverlay>
          <div className="relative">
            <div className={`flex touch-none items-center rounded px-1.5 py-1 ${getColorClassName(activeSchedule.color)}`}>
              <span className="line-clamp-2 text-[0.6rem] md:line-clamp-1 md:text-xs">{generateDragLabel()}</span>
            </div>
            {bulkCount > 1 && (
              <span className="absolute -left-2 -top-2 flex min-w-max items-center justify-center rounded-full bg-blue-600 px-1.5 py-0.5 text-[0.6rem] font-bold text-white shadow">
                {bulkCount}件移動中
              </span>
            )}
          </div>
        </DragOverlay>
      )}
    </DndContext>
  );
};
