import { useState, useEffect } from 'react';
import { isBefore, startOfDay } from 'date-fns';

import { BaseDialog } from '@/component/dialog/BaseDialog';
import { ErrorMessage } from '@/component/dialog/ErrorMessage';
import { ScheduleTypeButton } from '../ScheduleTypeButton';
import { SelectColor } from '../SelectColor';
import { InputScheduleName } from '@/component/dialog/InputScheduleName';
import { MasterScheduleTemplates } from './MasterScheduleTemplates';
import { CustomScheduleSelects } from './CustomScheduleSelects';
import { OptionCustomScheduleBulkCreate } from './OptionCustomScheduleBulkCreate';
import { InputDuration } from '@/component/dialog/InputDuration';
import { CreateButton } from './CreateButton';
import { CancelButton } from '../CancelButton';

import { getFirstColorKey } from '@/component/calendar/color-module';
import { toScheduleTypeName } from '@/component/schedule/schedule-module';
import { CreateSchedule } from '@/model/createSchedule';

type Props = {
  defaultType?: 'master' | 'custom';
  defaultDate?: Date | null;
  isOpen: boolean;
  submit: (createSchedule: CreateSchedule) => Promise<void>;
  close: () => void;
};

export const CreateScheduleDialog = ({ defaultType, defaultDate, isOpen, close, submit }: Props) => {
  const [name, setName] = useState('');
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date());
  const [scheduleType, setScheduleType] = useState<'master' | 'custom'>('custom');
  const [colorKey, setColorKey] = useState(getFirstColorKey());
  const [hasBulk, setHasBulk] = useState(false);
  const [bulkFrom, setBulkFrom] = useState(1);
  const [bulkTo, setBulkTo] = useState(8);
  const [errorMessage, setErrorMessage] = useState('');
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setName('');
    setStartDate(defaultDate ?? new Date());
    setEndDate(defaultDate ?? new Date());
    setScheduleType(defaultType ?? 'custom');
    setColorKey(getFirstColorKey());
    setHasBulk(false);
    setBulkFrom(1);
    setBulkTo(8);
    setErrorMessage('');
  }, [defaultDate, defaultType, isOpen]);

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

  const create = async () => {
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

    setLoading(true);
    await submit(createSchedule);
    close();
    setLoading(false);
  };

  return (
    <BaseDialog isOpen={isOpen} onClose={close} title="スケジュールを追加" disabled={loading}>
      {errorMessage && <ErrorMessage message={errorMessage} />}
      <div className="flex gap-2">
        <ScheduleTypeButton label={toScheduleTypeName('custom')} isSelected={scheduleType === 'custom'} onClick={onSelectCustom} />
        <ScheduleTypeButton label={toScheduleTypeName('master')} isSelected={scheduleType === 'master'} onClick={onSelectMaster} />
        <SelectColor value={colorKey} onChange={setColorKey} />
      </div>
      <InputScheduleName value={name} onChange={setName} />
      {scheduleType === 'master' && <MasterScheduleTemplates onSelect={onSelectMasterScheduleTemplate} />}
      {scheduleType === 'custom' && (
        <>
          <CustomScheduleSelects onSelect={(name: string, color: string) => (setName(name), setColorKey(color))} />
          <OptionCustomScheduleBulkCreate
            checked={hasBulk}
            setChecked={setHasBulk}
            from={bulkFrom}
            setFrom={setBulkFrom}
            to={bulkTo}
            setTo={setBulkTo}
          />
        </>
      )}
      <InputDuration from={startDate} to={endDate} onChangeFrom={setStartDate} onChangeTo={setEndDate} />
      <div className="flex justify-end gap-2">
        <CreateButton onClick={create} loading={loading} />
        <CancelButton onClick={close} disabled={loading} />
      </div>
    </BaseDialog>
  );
};
