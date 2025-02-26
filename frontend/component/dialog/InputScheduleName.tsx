type Props = {
  value: string;
  onChange: (value: string) => void;
};

export const InputScheduleName = ({ value, onChange }: Props) => {
  return (
    <input
      type="text"
      className="w-full border-b py-1 focus-visible:border-blue-500 focus-visible:outline-none"
      placeholder="スケジュール名"
      value={value}
      onChange={(e) => onChange(e.target.value)}
    />
  );
};
