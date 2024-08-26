import { useDraggable } from '@dnd-kit/core';
import { CSS } from '@dnd-kit/utilities';

import { ScheduleItem } from '@/component/schedule/ScheduleItem';

import { EditSchedule } from '@/model/edit-schedule';
import type { Schedule } from '@/type/schedule';

type Props = {
  schedule: Schedule;
  removeSchedule: (id: string) => void;
  saveSchedule: (editSchedule: EditSchedule) => void;
  changeScheduleColor: (id: string, color: string) => void;
};

export const ScheduleWeekItem = ({ schedule, removeSchedule, saveSchedule, changeScheduleColor }: Props) => {
  const { attributes, listeners, setNodeRef, transform } = useDraggable({ id: schedule.id });
  const style = {
    transform: CSS.Translate.toString(transform),
  };

  return (
    <div ref={setNodeRef} style={style} {...attributes} {...listeners}>
      <ScheduleItem
        schedule={schedule}
        removeSchedule={removeSchedule}
        saveSchedule={saveSchedule}
        changeScheduleColor={changeScheduleColor}
      />
    </div>
  );
};
