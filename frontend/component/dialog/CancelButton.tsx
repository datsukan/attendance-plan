type Props = {
  onClick: () => void;
};

export const CancelButton = ({ onClick }: Props) => {
  return (
    <button className="rounded-md px-3 py-1 hover:bg-gray-100 active:bg-gray-200" onClick={onClick}>
      <span>キャンセル</span>
    </button>
  );
};
