import { useState, useEffect } from 'react';
import { isBefore, startOfDay } from 'date-fns';

import { BaseDialog } from '@/component/schedule-dialog/BaseDialog';
import { ErrorMessage } from '@/component/schedule-dialog/ErrorMessage';
import { ScheduleTypeButton } from '../schedule-dialog/ScheduleTypeButton';
import { SelectColor } from '../schedule-dialog/SelectColor';
import { InputScheduleName } from '@/component/schedule-dialog/InputScheduleName';
import { MasterScheduleTemplates } from './MasterScheduleTemplates';
import { OptionCustomScheduleBulkCreate } from './OptionCustomScheduleBulkCreate';
import { InputDuration } from '@/component/schedule-dialog/InputDuration';
import { CreateButton } from './CreateButton';
import { CancelButton } from '../schedule-dialog/CancelButton';

import { getFirstColorKey } from '@/component/calendar/color-module';
import { toScheduleTypeName } from '@/component/schedule/schedule-module';
import { CreateSchedule } from '@/model/create-schedule';

type Props = {
  isOpen: boolean;
  submit: (createSchedule: CreateSchedule) => void;
  close: () => void;
};

export const AddScheduleDialog = ({ isOpen, close, submit }: Props) => {
  const [name, setName] = useState('');
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date());
  const [scheduleType, setScheduleType] = useState<'master' | 'custom'>('custom');
  const [colorKey, setColorKey] = useState(getFirstColorKey());
  const [hasBulk, setHasBulk] = useState(false);
  const [bulkFrom, setBulkFrom] = useState(1);
  const [bulkTo, setBulkTo] = useState(10);
  const [errorMessage, setErrorMessage] = useState('');

  useEffect(() => {
    setName('');
    setStartDate(new Date());
    setEndDate(new Date());
    setScheduleType('custom');
    setColorKey(getFirstColorKey());
    setHasBulk(false);
    setBulkFrom(1);
    setBulkTo(10);
    setErrorMessage('');
  }, [isOpen]);

  const onSelectCustom = () => {
    setScheduleType('custom');
    setColorKey('white');
  };

  const onSelectMaster = () => {
    setScheduleType('master');
    setColorKey('orange');
  };

  const onSelectMasterScheduleTemplate = (name: string, color: string) => {
    setName(name);
    setColorKey(color);
  };

  const create = () => {
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

    const createSchedule = new CreateSchedule(name, mStartDate, mEndDate, colorKey, scheduleType, hasBulk);
    if (hasBulk) {
      createSchedule.setBulkFrom(bulkFrom);
      createSchedule.setBulkTo(bulkTo);
    }
    submit(createSchedule);
    close();
  };

  return (
    <BaseDialog isOpen={isOpen} onClose={close} title="スケジュールを追加">
      {errorMessage && <ErrorMessage message={errorMessage} />}
      <div className="flex gap-2">
        <ScheduleTypeButton label={toScheduleTypeName('custom')} isSelected={scheduleType === 'custom'} onClick={onSelectCustom} />
        <ScheduleTypeButton label={toScheduleTypeName('master')} isSelected={scheduleType === 'master'} onClick={onSelectMaster} />
        <SelectColor value={colorKey} onChange={setColorKey} />
      </div>
      <InputScheduleName value={name} onChange={setName} />
      {scheduleType === 'master' && <MasterScheduleTemplates onSelect={onSelectMasterScheduleTemplate} />}
      {scheduleType === 'custom' && (
        <OptionCustomScheduleBulkCreate
          checked={hasBulk}
          setChecked={setHasBulk}
          from={bulkFrom}
          setFrom={setBulkFrom}
          to={bulkTo}
          setTo={setBulkTo}
        />
      )}
      <InputDuration from={startDate} to={endDate} onChangeFrom={setStartDate} onChangeTo={setEndDate} />
      <div className="flex gap-2 justify-end">
        <CreateButton onClick={create} />
        <CancelButton onClick={close} />
      </div>
    </BaseDialog>
  );
};
