import { ExclamationCircleIcon } from '@heroicons/react/24/outline';

type Props = {
  message: string;
};

export const ErrorMessage = ({ message }: Props) => {
  return (
    <div className="text-red-500 bg-red-100 rounded-lg p-2 flex gap-2 items-center">
      <ExclamationCircleIcon className="size-5" />
      <span className="text-sm">{message}</span>
    </div>
  );
};
