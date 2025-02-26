type Props = {
  name: string;
  label: string;
  type: string;
  autocomplete?: string;
  defaultValue?: string;
  errorMessage: string;
  disabled?: boolean;
};

export const InputTextField = ({ name, label, type, autocomplete = '', defaultValue, errorMessage, disabled = false }: Props) => {
  return (
    <div className="flex w-full flex-col gap-2">
      <div className="flex content-center justify-between">
        <label className="text-xs" htmlFor={name}>
          {label}
        </label>
        <span className="line-clamp-1 max-w-[70%] text-xs font-bold text-red-500" title={errorMessage}>
          {errorMessage}
        </span>
      </div>
      <div>
        <input
          id={name}
          name={name}
          type={type}
          autoComplete={autocomplete}
          defaultValue={defaultValue}
          disabled={disabled}
          className="w-full rounded-lg border p-2"
        />
      </div>
    </div>
  );
};
