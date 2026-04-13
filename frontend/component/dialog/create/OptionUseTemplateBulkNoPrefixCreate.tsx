type Props = {
  checked: boolean;
  setChecked: (checked: boolean) => void;
};

export const OptionUseTemplateBulkNoPrefixCreate = ({ checked, setChecked }: Props) => {
  return (
    <div className="flex items-center gap-2">
      <input
        id="option-use-template-bulk-no-prefix-create"
        type="checkbox"
        className="size-5"
        checked={checked}
        onChange={() => setChecked(!checked)}
      />
      <label htmlFor="option-use-template-bulk-no-prefix-create" className={`cursor-pointer text-sm ${checked ? '' : 'text-gray-400'}`}>
        全テンプレート分を一括作成する
      </label>
    </div>
  );
};
