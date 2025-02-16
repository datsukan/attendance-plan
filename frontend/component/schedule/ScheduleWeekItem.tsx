import { useSortable } from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';

import { ScheduleItem } from '@/component/schedule/ScheduleItem';

import { Model } from '@/model';
import { Type } from '@/type';

type Props = {
  schedule: Type.Schedule;
  isActive: boolean;
  removeSchedule: (id: string, type: Type.ScheduleType) => void;
  saveSchedule: (editSchedule: Model.EditSchedule) => void;
  changeScheduleColor: (id: string, type: Type.ScheduleType, color: string) => void;
};

export const ScheduleWeekItem = ({ schedule, isActive, removeSchedule, saveSchedule, changeScheduleColor }: Props) => {
  const { attributes, listeners, setNodeRef, transform, transition } = useSortable({ id: schedule.id, data: { date: schedule.startDate } });
  const style = {
    transform: transform ? CSS.Transform.toString(transform) : undefined,
    transition,
  };

  return (
    <div ref={setNodeRef} style={style} {...attributes} {...listeners} className={isActive ? 'opacity-50' : ''}>
      <ScheduleItem
        schedule={schedule}
        removeSchedule={removeSchedule}
        saveSchedule={saveSchedule}
        changeScheduleColor={changeScheduleColor}
      />
    </div>
  );
};
