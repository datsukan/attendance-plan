import { useState, useEffect } from 'react';
import { isBefore, startOfDay } from 'date-fns';

import { Type } from '@/type';
import { ScheduleTypeCustom } from '@/const/schedule';

import { BaseDialog } from '@/component/dialog/BaseDialog';
import { ErrorMessage } from '@/component/dialog/ErrorMessage';
import { ScheduleTypeButton } from '../ScheduleTypeButton';
import { SelectColor } from '../SelectColor';
import { InputScheduleName } from '@/component/dialog/InputScheduleName';
import { MasterScheduleTemplates } from './MasterScheduleTemplates';
import { CustomScheduleSelects } from './CustomScheduleSelects';
import { OptionRangeBulkCreate } from './OptionRangeBulkCreate';
import { OptionUseTemplateBulkCreate } from './OptionUseTemplateBulkCreate';
import { InputDuration } from '@/component/dialog/InputDuration';
import { CreateButton } from './CreateButton';
import { CancelButton } from '../CancelButton';

import { getFirstColorKey } from '@/component/calendar/color-module';
import { toScheduleTypeName } from '@/component/schedule/schedule-module';
import { CreateSchedule } from '@/model/createSchedule';
import { useSubject } from '@/provider/SubjectProvider';

type Props = {
  defaultType?: Type.ScheduleType;
  defaultDate?: Date | null;
  isOpen: boolean;
  submit: (createSchedule: CreateSchedule[]) => Promise<void>;
  close: () => void;
};

export const CreateScheduleDialog = ({ defaultType, defaultDate, isOpen, close, submit }: Props) => {
  const [name, setName] = useState('');
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date());
  const [scheduleType, setScheduleType] = useState<'master' | 'custom'>('custom');
  const [colorKey, setColorKey] = useState(getFirstColorKey());
  const [hasRangeBulk, setHasRangeBulk] = useState(false);
  const [rangeBulkFrom, setRangeBulkFrom] = useState(1);
  const [rangeBulkTo, setRangeBulkTo] = useState(8);
  const [hasUseTempBulk, setHasUseTempBulk] = useState(false);
  const [useTempBulkNumber, setUseTempBulkNumber] = useState(1);
  const [errorMessage, setErrorMessage] = useState('');
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setName('');
    setStartDate(defaultDate ?? new Date());
    setEndDate(defaultDate ?? new Date());
    setScheduleType(defaultType ?? 'custom');
    setColorKey(getFirstColorKey());
    setHasRangeBulk(false);
    setRangeBulkFrom(1);
    setRangeBulkTo(8);
    setHasUseTempBulk(false);
    setUseTempBulkNumber(1);
    setErrorMessage('');
  }, [defaultDate, defaultType, isOpen]);

  const { subjects } = useSubject();

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

  const validateNormal = (): boolean => {
    let result = true;

    if (!name) {
      setErrorMessage('スケジュール名を入力してください');
      result = false;
    }

    if (isBefore(startOfDay(endDate), startOfDay(startDate))) {
      setErrorMessage('終了日は開始日以降にしてください');
      result = false;
    }

    return result;
  };

  const validateUseTempBulk = (): boolean => {
    if (isBefore(startOfDay(endDate), startOfDay(startDate))) {
      setErrorMessage('終了日は開始日以降にしてください');
      return false;
    }

    return true;
  };

  const genSchedulesNormal = (): CreateSchedule[] => {
    return [new CreateSchedule(name, startOfDay(startDate), startOfDay(endDate), colorKey, scheduleType)];
  };

  const genSchedulesByRangeBulk = (): CreateSchedule[] => {
    const schedules: CreateSchedule[] = [];
    for (let n = rangeBulkFrom; n <= rangeBulkTo; n++) {
      schedules.push(new CreateSchedule(`第${n}回 ${name}`, startOfDay(startDate), startOfDay(endDate), colorKey, scheduleType));
    }
    return schedules;
  };

  const genSchedulesByUseTempBulk = (): CreateSchedule[] => {
    const schedules: CreateSchedule[] = [];
    for (const subject of subjects) {
      schedules.push(
        new CreateSchedule(
          `第${useTempBulkNumber}回 ${subject.name}`,
          startOfDay(startDate),
          startOfDay(endDate),
          subject.color,
          scheduleType
        )
      );
    }
    return schedules;
  };

  const validate = (): boolean => {
    if (scheduleType === ScheduleTypeCustom && hasUseTempBulk) {
      return validateUseTempBulk();
    }

    return validateNormal();
  };

  const genSchedules = (): CreateSchedule[] => {
    if (scheduleType === ScheduleTypeCustom) {
      if (hasRangeBulk) {
        return genSchedulesByRangeBulk();
      } else if (hasUseTempBulk) {
        return genSchedulesByUseTempBulk();
      }
    }

    return genSchedulesNormal();
  };

  const create = async () => {
    if (!validate()) {
      return;
    }

    const schedules = genSchedules();
    if (schedules.length === 0) {
      return;
    }

    setLoading(true);
    await submit(schedules);
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
      {!hasUseTempBulk && <InputScheduleName value={name} onChange={setName} />}
      {scheduleType === 'master' && <MasterScheduleTemplates onSelect={onSelectMasterScheduleTemplate} />}
      {scheduleType === 'custom' && (
        <>
          <CustomScheduleSelects onSelect={(name: string, color: string) => (setName(name), setColorKey(color))} />
          {!hasUseTempBulk && (
            <OptionRangeBulkCreate
              checked={hasRangeBulk}
              setChecked={setHasRangeBulk}
              from={rangeBulkFrom}
              setFrom={setRangeBulkFrom}
              to={rangeBulkTo}
              setTo={setRangeBulkTo}
            />
          )}
          {!hasRangeBulk && (
            <OptionUseTemplateBulkCreate
              number={useTempBulkNumber}
              checked={hasUseTempBulk}
              setNumber={setUseTempBulkNumber}
              setChecked={setHasUseTempBulk}
            />
          )}
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
