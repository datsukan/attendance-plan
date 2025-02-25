import { useSortable } from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';

import { ScheduleItem } from '@/component/schedule/ScheduleItem';

import { Model } from '@/model';
import { Type } from '@/type';

type Props = {
  schedule: Type.Schedule;
  isActive: boolean;
};

export const ScheduleWeekItem = ({ schedule, isActive }: Props) => {
  const { attributes, listeners, setNodeRef, transform, transition } = useSortable({ id: schedule.id, data: { date: schedule.startDate } });
  const style = {
    transform: transform ? CSS.Transform.toString(transform) : undefined,
    transition,
  };

  return (
    <div ref={setNodeRef} style={style} {...attributes} {...listeners} className={isActive ? 'opacity-50' : ''}>
      <ScheduleItem schedule={schedule} />
    </div>
  );
};
