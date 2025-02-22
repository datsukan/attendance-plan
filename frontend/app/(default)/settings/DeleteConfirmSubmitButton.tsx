import { ArrowPathIcon } from '@heroicons/react/24/outline';

type Props = {
  loading: boolean;
  onClick: () => void;
};

export const DeleteConfirmSubmitButton = ({ loading, onClick }: Props) => {
  return (
    <button
      className="rounded-md bg-red-600 px-3 py-1 hover:bg-red-500 active:bg-red-400 disabled:bg-red-200"
      onClick={onClick}
      disabled={loading}
    >
      <span className="text-white">{loading ? <ArrowPathIcon className="size-5 animate-spin" /> : '削除'}</span>
    </button>
  );
};
