import { InputDate } from '@/component/dialog/InputDate';

type Props = {
  from: Date;
  to: Date;
  onChangeFrom: (value: Date) => void;
  onChangeTo: (value: Date) => void;
};

export const InputDuration = ({ from, to, onChangeFrom, onChangeTo }: Props) => {
  const handleFromBlur = () => {
    if (from > to) {
      onChangeTo(from);
    }
  };

  const handleToBlur = () => {
    if (to < from) {
      onChangeFrom(to);
    }
  };

  return (
    <div className="flex w-full items-center gap-2">
      <InputDate value={from} onChange={onChangeFrom} onBlur={handleFromBlur} />
      <span>-</span>
      <InputDate value={to} onChange={onChangeTo} onBlur={handleToBlur} />
    </div>
  );
};
