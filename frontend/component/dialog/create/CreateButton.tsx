type Props = {
  onClick: () => void;
};

export const CreateButton = ({ onClick }: Props) => {
  return (
    <button className="rounded-md px-3 py-1 bg-blue-600 hover:bg-blue-500 active:bg-blue-400" onClick={onClick}>
      <span className="text-white">作成</span>
    </button>
  );
};
