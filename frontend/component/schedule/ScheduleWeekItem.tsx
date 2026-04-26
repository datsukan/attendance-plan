import type { CSSProperties } from 'react';
import { useSortable } from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';

import { ScheduleItem } from '@/component/schedule/ScheduleItem';
import { useSelection } from '@/provider/SelectionContext';

import { Type } from '@/type';

type Props = {
  schedule: Type.Schedule;
  isActive: boolean;
  colStartClassName: string;
  colEndClassName: string;
};

export const ScheduleWeekItem = ({ schedule, isActive, colStartClassName, colEndClassName }: Props) => {
  const { isSelected } = useSelection();
  const { attributes, listeners, setNodeRef, transform, transition } = useSortable({
    id: schedule.id,
    data: { date: schedule.startDate, type: schedule.type },
  });

  const style: CSSProperties = {
    transform: CSS.Transform.toString(transform),
    transition,
  };

  const selected = isSelected(schedule.id);

  return (
    <div
      ref={setNodeRef}
      style={style}
      {...attributes}
      {...listeners}
      tabIndex={-1}
      className={[
        'pointer-events-auto mr-2',
        colStartClassName,
        colEndClassName,
        isActive ? 'opacity-50' : '',
        selected ? 'rounded ring-2 ring-blue-500 ring-offset-1' : '',
      ]
        .filter(Boolean)
        .join(' ')}
    >
      <ScheduleItem schedule={schedule} />
    </div>
  );
};
