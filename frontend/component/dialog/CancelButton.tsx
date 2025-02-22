type Props = {
  disabled?: boolean;
  onClick: () => void;
};

export const CancelButton = ({ disabled = false, onClick }: Props) => {
  return (
    <button className="rounded-md px-3 py-1 hover:bg-gray-100 active:bg-gray-200 disabled:bg-white" onClick={onClick} disabled={disabled}>
      <span>キャンセル</span>
    </button>
  );
};
