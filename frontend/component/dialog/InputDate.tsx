import { format } from 'date-fns';

type Props = {
  value: Date;
  onChange: (value: Date) => void;
};

export const InputDate = ({ value, onChange }: Props) => {
  return (
    <input
      type="date"
      className="border-b py-1 w-full focus-visible:outline-none focus-visible:border-blue-500"
      value={format(value, 'yyyy-MM-dd')}
      onChange={(e) => onChange(new Date(e.target.value))}
    />
  );
};
