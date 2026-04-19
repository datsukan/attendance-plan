'use client';

import { useState } from 'react';
import toast from 'react-hot-toast';
import { CheckCircleIcon, XMarkIcon, ArrowUturnLeftIcon } from '@heroicons/react/24/outline';

import { SessionExpiredError } from '@/backend-api/error';

type Props = {
  toastId: string;
  label: string;
  onUndo: () => Promise<void>;
  visible: boolean;
};

export const UndoToast = ({ toastId, label, onUndo, visible }: Props) => {
  const [isExecuting, setIsExecuting] = useState(false);

  const handleUndo = async () => {
    setIsExecuting(true);
    try {
      await onUndo();
    } catch (e) {
      if (!(e instanceof SessionExpiredError)) {
        toast.error(String(e));
      }
    } finally {
      toast.dismiss(toastId);
      setIsExecuting(false);
    }
  };

  return (
    <div
      className={`flex flex-wrap items-center gap-x-2 gap-y-1 rounded-lg bg-white px-3 py-2 shadow-lg transition-opacity duration-300 ${
        visible ? 'opacity-100' : 'opacity-0'
      }`}
    >
      <div className="flex min-w-0 flex-1 items-center gap-2">
        <CheckCircleIcon className="size-5 shrink-0 text-green-500" />
        <span className="min-w-0 break-words text-sm text-gray-800">{label}</span>
      </div>
      <div className="ml-auto flex shrink-0 items-center gap-1">
        <button
          onClick={handleUndo}
          disabled={isExecuting}
          className="flex items-center gap-1 rounded px-2 py-1 text-sm font-medium text-blue-600 hover:bg-blue-50 disabled:opacity-50"
        >
          {isExecuting ? (
            <span className="size-4 animate-spin rounded-full border-2 border-blue-600 border-t-transparent" />
          ) : (
            <ArrowUturnLeftIcon className="size-4" />
          )}
          取り消す
        </button>
        <button
          onClick={() => toast.dismiss(toastId)}
          disabled={isExecuting}
          className="rounded-full p-1 hover:bg-gray-100 disabled:opacity-50"
        >
          <XMarkIcon className="size-5 text-gray-500" />
        </button>
      </div>
    </div>
  );
};
