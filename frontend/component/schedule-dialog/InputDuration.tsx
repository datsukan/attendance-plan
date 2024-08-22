import { InputDate } from '@/component/schedule-dialog/InputDate';

type Props = {
  from: Date;
  to: Date;
  onChangeFrom: (value: Date) => void;
  onChangeTo: (value: Date) => void;
};

export const InputDuration = ({ from, to, onChangeFrom, onChangeTo }: Props) => {
  return (
    <div className="w-full flex gap-2 items-center">
      <InputDate value={from} onChange={onChangeFrom} />
      <span>-</span>
      <InputDate value={to} onChange={onChangeTo} />
    </div>
  );
};
