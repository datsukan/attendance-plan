import { ArrowPathIcon } from '@heroicons/react/24/outline';

type Props = {
  label: string;
  loadingLabel?: string;
  loading?: boolean;
};

export const SubmitButton = ({ label, loadingLabel = '', loading = false }: Props) => {
  return (
    <button
      type="submit"
      className="w-full rounded-lg bg-blue-600 px-4 py-2 text-white hover:bg-blue-800 disabled:bg-blue-400"
      disabled={loading}
    >
      {loading ? (
        <div className="flex items-center justify-center gap-2">
          <ArrowPathIcon className="size-5 animate-spin" />
          {loadingLabel}
        </div>
      ) : (
        label
      )}
    </button>
  );
};
