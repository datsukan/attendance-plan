import { format, isEqual } from 'date-fns';
import { ja } from 'date-fns/locale';
import { TrashIcon, PencilIcon } from '@heroicons/react/24/outline';

import { Type } from '@/type';
import { getColorClassName, getColorKeys } from '@/component/calendar/color-module';
import { toScheduleTypeName } from '@/component/schedule/schedule-module';

type Props = {
  schedule: Type.Schedule;
  onSelectColor: (color: string) => void;
  openRemoveConfirmDialog: () => void;
  openEditDialog: () => void;
};

export const InfoCard = ({ schedule, onSelectColor, openRemoveConfirmDialog, openEditDialog }: Props) => {
  return (
    <div className="divide-y overflow-hidden rounded-lg bg-white shadow-lg">
      <div className="space-y-2 p-3">
        <div className="flex items-center gap-2">
          <div className="rounded border px-2 py-1 text-sm">{toScheduleTypeName(schedule.type)}</div>
          <div className={`size-5 rounded-full ${getColorClassName(schedule.color)}`}></div>
          <div className="text-lg">{schedule.name}</div>
        </div>
        <div className="text-sm">
          {format(schedule.startDate, 'yyyy年M月d日(E)', { locale: ja })}
          {isEqual(schedule.startDate, schedule.endDate) ? '' : ` ~ ${format(schedule.endDate, 'yyyy年M月d日(E)', { locale: ja })}`}
        </div>
      </div>
      <div>
        <button className="flex w-full items-center gap-2 px-4 py-2 hover:bg-gray-100 active:bg-gray-200" onClick={() => openEditDialog()}>
          <PencilIcon className="size-5 text-gray-600" />
          <span>編集</span>
        </button>
        <button className="flex w-full items-center gap-2 px-4 py-2 hover:bg-gray-100 active:bg-gray-200" onClick={openRemoveConfirmDialog}>
          <TrashIcon className="size-5 text-gray-600" />
          <span>削除</span>
        </button>
      </div>
      <div>
        <div className="grid w-fit grid-cols-4 gap-2 p-2">
          {getColorKeys().map((color) => (
            <button
              key={color}
              className={`mx-auto size-6 rounded-full ${getColorClassName(color)}`}
              onClick={() => onSelectColor(color)}
            />
          ))}
        </div>
      </div>
    </div>
  );
};
