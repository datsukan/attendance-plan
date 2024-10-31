type Props = {
  value: string;
  onChange: (value: string) => void;
};

export const InputScheduleName = ({ value, onChange }: Props) => {
  return (
    <input
      type="text"
      className="border-b w-full py-1 focus-visible:outline-none focus-visible:border-blue-500"
      placeholder="スケジュール名"
      value={value}
      onChange={(e) => onChange(e.target.value)}
    />
  );
};