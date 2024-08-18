import { useState } from 'react';
import { PlusIcon } from '@heroicons/react/24/outline';

import { AddScheduleDialog } from './AddScheduleDialog';

type Props = {
  create: (name: string, startDate: Date, endDate: Date) => void;
};

export const AddMasterScheduleButton = ({ create }: Props) => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <button className="rounded-full px-4 py-1.5 bg-blue-600 hover:bg-blue-500 active:bg-blue-400" onClick={() => setIsOpen(true)}>
        <div className="flex gap-1 items-center text-white">
          <PlusIcon className="size-5" />
          <span className="text-sm">作成</span>
        </div>
      </button>
      <AddScheduleDialog isOpen={isOpen} close={() => setIsOpen(false)} submit={create} />
    </>
  );
};
