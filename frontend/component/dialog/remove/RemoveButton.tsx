type Props = {
  onClick: () => void;
};

export const RemoveButton = ({ onClick }: Props) => {
  return (
    <button className="rounded-md bg-red-600 px-3 py-1 hover:bg-red-500 active:bg-red-400" onClick={onClick}>
      <span className="text-white">削除</span>
    </button>
  );
};
