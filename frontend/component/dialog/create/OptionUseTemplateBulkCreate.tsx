type Props = {
  number: number;
  checked: boolean;
  setNumber: (number: number) => void;
  setChecked: (checked: boolean) => void;
};

export const OptionUseTemplateBulkCreate = ({ number, checked, setNumber, setChecked }: Props) => {
  return (
    <div className="flex items-center gap-2">
      <input
        id="option-use-template-bulk-create"
        type="checkbox"
        className="size-5"
        checked={checked}
        onChange={() => setChecked(!checked)}
      />
      <label htmlFor="option-use-template-bulk-create" className={`cursor-pointer text-sm ${checked ? '' : 'text-gray-400'}`}>
        第
        <input
          type="number"
          min={1}
          max={99}
          className="w-10 border-b py-1 text-right focus-visible:border-blue-500 focus-visible:outline-none"
          value={number}
          onChange={(e) => setNumber(Number(e.target.value))}
          disabled={!checked}
        />
        回 で全テンプレート分を一括作成する
      </label>
    </div>
  );
};
