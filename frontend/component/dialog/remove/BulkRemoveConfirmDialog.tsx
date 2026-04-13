'use client';

import { useState } from 'react';
import { format, isEqual } from 'date-fns';
import { ja } from 'date-fns/locale';

import { BaseDialog } from '@/component/dialog/BaseDialog';
import { RemoveButton } from './RemoveButton';
import { CancelButton } from '@/component/dialog/CancelButton';

import { Type } from '@/type';
import { toScheduleTypeName } from '@/component/schedule/schedule-module';
import { getColorClassName } from '@/component/calendar/color-module';

type Props = {
  schedules: Type.Schedule[];
  isOpen: boolean;
  close: () => void;
  remove: () => Promise<void>;
};

export const BulkRemoveConfirmDialog = ({ schedules, isOpen, close, remove }: Props) => {
  const [loading, setLoading] = useState(false);

  const removeSchedules = async () => {
    setLoading(true);
    await remove();
    close();
    setLoading(false);
  };

  return (
    <BaseDialog isOpen={isOpen} onClose={close} title="スケジュールの一括削除" disabled={loading}>
      <div className="text-red-500">選択中の {schedules.length} 件のスケジュールを削除しますか？</div>
      <div className="max-h-64 space-y-2 overflow-y-auto">
        {schedules.map((schedule) => (
          <div key={schedule.id} className="space-y-1 rounded-lg border p-3">
            <div className="flex items-center gap-2">
              <div className="rounded border px-2 py-1 text-sm">{toScheduleTypeName(schedule.type)}</div>
              <div className={`size-5 rounded-full ${getColorClassName(schedule.color)}`}></div>
              <div className="text-base">{schedule.name}</div>
            </div>
            <div className="text-sm text-gray-500">
              {format(schedule.startDate, 'yyyy年M月d日(E)', { locale: ja })}
              {isEqual(schedule.startDate, schedule.endDate) ? '' : ` ~ ${format(schedule.endDate, 'yyyy年M月d日(E)', { locale: ja })}`}
            </div>
          </div>
        ))}
      </div>
      <div className="flex justify-end gap-2">
        <RemoveButton onClick={removeSchedules} loading={loading} />
        <CancelButton onClick={close} disabled={loading} />
      </div>
    </BaseDialog>
  );
};
