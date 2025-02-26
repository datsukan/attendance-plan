import { useState } from 'react';
import { PlusIcon } from '@heroicons/react/24/outline';

import { CreateScheduleDialog } from '@/component/dialog/create/CreateScheduleDialog';

import { CreateSchedule } from '@/model/createSchedule';

type Props = {
  create: (createSchedule: CreateSchedule) => Promise<void>;
};

export const AddScheduleButton = ({ create }: Props) => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <button className="rounded-full bg-blue-600 py-1.5 pl-3 pr-4 hover:bg-blue-500 active:bg-blue-400" onClick={() => setIsOpen(true)}>
        <div className="flex items-center gap-1 text-white">
          <PlusIcon className="size-5" />
          <span className="mb-0.5 text-sm">作成</span>
        </div>
      </button>
      <CreateScheduleDialog isOpen={isOpen} close={() => setIsOpen(false)} submit={create} />
    </>
  );
};
