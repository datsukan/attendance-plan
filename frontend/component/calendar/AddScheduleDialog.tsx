import { useState, useEffect } from 'react';
import { Dialog, DialogBackdrop, DialogPanel, DialogTitle } from '@headlessui/react';
import { ExclamationCircleIcon } from '@heroicons/react/24/outline';
import { format, isBefore, startOfDay } from 'date-fns';

type Props = {
  isOpen: boolean;
  submit: (name: string, startDate: Date, endDate: Date) => void;
  close: () => void;
};

export const AddScheduleDialog = ({ isOpen, close, submit }: Props) => {
  const [name, setName] = useState('');
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date());
  const [errorMessage, setErrorMessage] = useState('');

  useEffect(() => {
    setName('');
    setStartDate(new Date());
    setEndDate(new Date());
    setErrorMessage('');
  }, [isOpen]);

  const create = () => {
    if (!name) {
      setErrorMessage('スケジュール名を入力してください');
      return;
    }

    if (isBefore(endDate, startDate)) {
      setErrorMessage('終了日は開始日以降にしてください');
      return;
    }

    submit(name, startDate, endDate);
    close();
  };

  return (
    <Dialog as="div" open={isOpen} onClose={close} className="relative z-50">
      <DialogBackdrop className="fixed inset-0 bg-black/30" />
      <div className="fixed inset-0 flex w-screen items-center justify-center p-4">
        <DialogPanel className="max-w-lg min-w-96 space-y-4 border rounded-xl bg-white p-6">
          <DialogTitle className="text-lg font-semibold">スケジュールを追加</DialogTitle>
          {errorMessage && (
            <div className="text-red-500 bg-red-100 rounded-lg p-2 flex gap-2 items-center">
              <ExclamationCircleIcon className="size-5" />
              <span className="text-sm">{errorMessage}</span>
            </div>
          )}
          <input
            type="text"
            className="border-b w-full py-1"
            placeholder="スケジュール名"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <div className="w-full flex gap-2 items-center">
            <input
              type="date"
              className="border-b py-1 w-full"
              value={format(startDate, 'yyyy-MM-dd')}
              onChange={(e) => setStartDate(startOfDay(new Date(e.target.value)))}
            />
            <span>-</span>
            <input
              type="date"
              className="border-b py-1 w-full"
              value={format(endDate, 'yyyy-MM-dd')}
              onChange={(e) => setEndDate(startOfDay(new Date(e.target.value)))}
            />
          </div>
          <div className="flex gap-2 justify-end">
            <button className="rounded-md px-3 py-1 bg-blue-600" onClick={() => create()}>
              <span className="text-white">作成</span>
            </button>
            <button className="border rounded-md px-3 py-1" onClick={close}>
              <span>キャンセル</span>
            </button>
          </div>
        </DialogPanel>
      </div>
    </Dialog>
  );
};
