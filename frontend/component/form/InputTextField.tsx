type Props = {
  name: string;
  label: string;
  type: string;
  errorMessage: string;
};

export const InputTextField = ({ name, label, type, errorMessage }: Props) => {
  return (
    <div className="flex w-full flex-col gap-2">
      <div className="flex content-center justify-between">
        <label className="text-xs" htmlFor={name}>
          {label}
        </label>
        <span className="line-clamp-1 max-w-[70%] text-xs font-bold text-red-500">{errorMessage}</span>
      </div>
      <div>
        <input id={name} name={name} type={type} className="w-full rounded-lg border p-2" />
      </div>
    </div>
  );
};
