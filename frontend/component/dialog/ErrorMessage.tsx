import { ExclamationCircleIcon } from '@heroicons/react/24/outline';

type Props = {
  message: string;
};

export const ErrorMessage = ({ message }: Props) => {
  return (
    <div className="flex items-center gap-2 rounded-lg bg-red-100 p-2 text-red-500">
      <ExclamationCircleIcon className="size-5" />
      <span className="text-sm">{message}</span>
    </div>
  );
};
