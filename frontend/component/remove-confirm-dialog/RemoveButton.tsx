type Props = {
  onClick: () => void;
};

export const RemoveButton = ({ onClick }: Props) => {
  return (
    <button className="rounded-md px-3 py-1 bg-red-600 hover:bg-red-500 active:bg-red-400" onClick={onClick}>
      <span className="text-white">削除</span>
    </button>
  );
};
