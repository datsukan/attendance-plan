import { ArrowPathIcon } from '@heroicons/react/24/outline';

type Props = {
  loading: boolean;
  onClick: () => void;
};

export const CreateButton = ({ loading, onClick }: Props) => {
  return (
    <button
      className="rounded-md bg-blue-600 px-3 py-1 hover:bg-blue-500 active:bg-blue-400 disabled:bg-blue-200"
      onClick={onClick}
      disabled={loading}
    >
      <span className="text-white">
        {loading ? (
          <div className="flex items-center justify-center gap-2">
            <ArrowPathIcon className="size-5 animate-spin" />
            作成中
          </div>
        ) : (
          '作成'
        )}
      </span>
    </button>
  );
};
