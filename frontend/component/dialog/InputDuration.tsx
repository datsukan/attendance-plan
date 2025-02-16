import { InputDate } from '@/component/dialog/InputDate';
import { useEffect } from 'react';

type Props = {
  from: Date;
  to: Date;
  onChangeFrom: (value: Date) => void;
  onChangeTo: (value: Date) => void;
};

export const InputDuration = ({ from, to, onChangeFrom, onChangeTo }: Props) => {
  useEffect(() => {
    if (from > to) {
      onChangeTo(from);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [from]);

  useEffect(() => {
    if (to < from) {
      onChangeFrom(to);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [to]);

  return (
    <div className="w-full flex gap-2 items-center">
      <InputDate value={from} onChange={onChangeFrom} />
      <span>-</span>
      <InputDate value={to} onChange={onChangeTo} />
    </div>
  );
};
