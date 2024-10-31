import { useSortable } from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';

import { ScheduleItem } from '@/component/schedule/ScheduleItem';

import { EditSchedule } from '@/model/edit-schedule';
import type { Schedule } from '@/type/schedule';

type Props = {
  schedule: Schedule;
  isActive: boolean;
  removeSchedule: (id: string) => void;
  saveSchedule: (editSchedule: EditSchedule) => void;
  changeScheduleColor: (id: string, color: string) => void;
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
