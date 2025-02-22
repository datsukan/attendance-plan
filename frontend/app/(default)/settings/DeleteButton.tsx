'use client';

import { useState } from 'react';

import { DeleteConfirmDialog } from './DeleteConfirmDialog';

export const DeleteButton = () => {
  const [isOpen, setIsOpen] = useState(false);

  const openDialog = () => setIsOpen(true);

  return (
    <>
      <button className="w-full rounded-lg bg-red-600 px-4 py-2 text-white hover:bg-red-800 disabled:bg-red-400" onClick={openDialog}>
        アカウントを削除する
      </button>
      <DeleteConfirmDialog isOpen={isOpen} close={() => setIsOpen(false)} />
    </>
  );
};
