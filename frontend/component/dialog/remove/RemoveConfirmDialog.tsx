import { format, isEqual } from 'date-fns';
import { ja } from 'date-fns/locale';

import { BaseDialog } from '@/component/dialog/BaseDialog';
import { RemoveButton } from './RemoveButton';
import { CancelButton } from '@/component/dialog/CancelButton';

import { Type } from '@/type';
import { toScheduleTypeName } from '@/component/schedule/schedule-module';
import { getColorClassName } from '@/component/calendar/color-module';

type Props = {
  schedule: Type.Schedule;
  isOpen: boolean;
  close: () => void;
  remove: () => void;
};

export const RemoveConfirmDialog = ({ schedule, isOpen, close, remove }: Props) => {
  const removeSchedule = () => {
    remove();
    close();
  };

  return (
    <BaseDialog isOpen={isOpen} onClose={close} title="スケジュールの削除">
      <div className="text-red-500">本当にこのスケジュールを削除しますか？</div>
      <div className="space-y-2 rounded-lg border p-3">
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
      <div className="flex justify-end gap-2">
        <RemoveButton onClick={removeSchedule} />
        <CancelButton onClick={close} />
      </div>
    </BaseDialog>
  );
};
