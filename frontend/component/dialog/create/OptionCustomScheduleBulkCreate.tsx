type Props = {
  checked: boolean;
  setChecked: (checked: boolean) => void;
  from: number;
  setFrom: (from: number) => void;
  to: number;
  setTo: (to: number) => void;
};

export const OptionCustomScheduleBulkCreate = ({ checked, setChecked, from, setFrom, to, setTo }: Props) => {
  return (
    <div className="flex items-center gap-2">
      <input
        id="option-custom-schedule-bulk-create"
        type="checkbox"
        className="size-5"
        checked={checked}
        onChange={() => setChecked(!checked)}
      />
      <label htmlFor="option-custom-schedule-bulk-create" className={`cursor-pointer text-sm ${checked ? '' : 'text-gray-400'}`}>
        <input
          type="number"
          min={1}
          max={to}
          className="border-b w-10 py-1 focus-visible:outline-none focus-visible:border-blue-500 text-right"
          value={from}
          onChange={(e) => setFrom(Number(e.target.value))}
          disabled={!checked}
        />
        回 から
        <input
          min={from}
          max={99}
          type="number"
          className="border-b w-10 py-1 focus-visible:outline-none focus-visible:border-blue-500 text-right"
          value={to}
          onChange={(e) => setTo(Number(e.target.value))}
          disabled={!checked}
        />
        回 で一括作成する
      </label>
    </div>
  );
};
