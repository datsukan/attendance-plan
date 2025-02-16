import { useEffect, useState } from 'react';
import { isBefore, startOfDay } from 'date-fns';

import { BaseDialog } from '@/component/dialog/BaseDialog';
import { ErrorMessage } from '@/component/dialog/ErrorMessage';
import { ScheduleTypeButton } from '@/component/dialog/ScheduleTypeButton';
import { SelectColor } from '@/component/dialog/SelectColor';
import { InputScheduleName } from '@/component/dialog/InputScheduleName';
import { InputDuration } from '@/component/dialog/InputDuration';
import { SaveButton } from '@/component/dialog/edit/SaveButton';
import { CancelButton } from '@/component/dialog/CancelButton';

import { toScheduleTypeName } from '@/component/schedule/schedule-module';
import { EditSchedule } from '@/model/editSchedule';
import { Type } from '@/type';

type Props = {
  schedule: Type.Schedule;
  isOpen: boolean;
  submit: (editSchedule: EditSchedule) => void;
  close: () => void;
};

export const EditScheduleDialog = ({ schedule, isOpen, submit, close }: Props) => {
  const [name, setName] = useState(schedule.name);
  const [startDate, setStartDate] = useState(schedule.startDate);
  const [endDate, setEndDate] = useState(schedule.endDate);
  const [colorKey, setColorKey] = useState(schedule.color);
  const [scheduleType, setScheduleType] = useState(schedule.type);
  const [errorMessage, setErrorMessage] = useState('');

  useEffect(() => {
    setName(schedule.name);
    setStartDate(schedule.startDate);
    setEndDate(schedule.endDate);
    setColorKey(schedule.color);
    setScheduleType(schedule.type);
    setErrorMessage('');
  }, [isOpen, schedule.name, schedule.startDate, schedule.endDate, schedule.color, schedule.type]);

  const save = () => {
    const mStartDate = startOfDay(startDate);
    const mEndDate = startOfDay(endDate);

    if (!name) {
      setErrorMessage('スケジュール名を入力してください');
      return;
    }

    if (isBefore(mEndDate, mStartDate)) {
      setErrorMessage('終了日は開始日以降にしてください');
      return;
    }

    const editSchedule = new EditSchedule(schedule);
    editSchedule.setName(name);
    editSchedule.setStartDate(mStartDate);
    editSchedule.setEndDate(mEndDate);
    editSchedule.setColor(colorKey);
    editSchedule.setType(scheduleType);

    submit(editSchedule);
    close();
  };

  return (
    <BaseDialog isOpen={isOpen} onClose={close} title="スケジュールを編集">
      {errorMessage && <ErrorMessage message={errorMessage} />}
      <div className="flex gap-2">
        <ScheduleTypeButton
          label={toScheduleTypeName('custom')}
          isSelected={scheduleType === 'custom'}
          onClick={() => setScheduleType('custom')}
        />
        <ScheduleTypeButton
          label={toScheduleTypeName('master')}
          isSelected={scheduleType === 'master'}
          onClick={() => setScheduleType('master')}
        />
        <SelectColor value={colorKey} onChange={setColorKey} />
      </div>
      <InputScheduleName value={name} onChange={setName} />
      <InputDuration from={startDate} to={endDate} onChangeFrom={setStartDate} onChangeTo={setEndDate} />
      <div className="flex justify-end gap-2">
        <SaveButton onClick={save} />
        <CancelButton onClick={close} />
      </div>
    </BaseDialog>
  );
};
