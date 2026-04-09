import { format } from 'date-fns';

type Props = {
  value: Date;
  onChange: (value: Date) => void;
  onBlur?: () => void;
};

export const InputDate = ({ value, onChange, onBlur }: Props) => {
  return (
    <input
      type="date"
      className="w-full border-b py-1 focus-visible:border-blue-500 focus-visible:outline-none"
      value={isNaN(value.getTime()) ? '' : format(value, 'yyyy-MM-dd')}
      onChange={(e) => {
        const d = new Date(e.target.value);
        if (!isNaN(d.getTime())) onChange(d);
      }}
      onBlur={onBlur}
    />
  );
};
