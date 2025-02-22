'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import toast from 'react-hot-toast';

import { BaseDialog } from '@/component/dialog/BaseDialog';
import { DeleteConfirmSubmitButton } from './DeleteConfirmSubmitButton';
import { CancelButton } from '@/component/dialog/CancelButton';

import { useStorage } from '@/provider/StorageProvider';
import { deleteUser } from '@/api/deleteUser';

type Props = {
  isOpen: boolean;
  close: () => void;
};

export const DeleteConfirmDialog = ({ isOpen, close }: Props) => {
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const { user, removeUser } = useStorage();

  if (!user) return null;

  const remove = () => {
    (async () => {
      setLoading(true);

      try {
        await deleteUser();
        removeUser();
        router.push('/signin');
        toast.success('アカウントを削除しました');
      } catch (e) {
        toast.error(String(e));
      }

      setLoading(false);
      close();
    })();
  };

  return (
    <BaseDialog isOpen={isOpen} onClose={close} title="アカウントの削除">
      <div className="text-red-500">本当にアカウントを削除しますか？</div>
      <div className="space-y-2 rounded-lg border p-3 text-sm">
        {user.name && <p>{user.name}</p>}
        <p>{user.email}</p>
      </div>
      <div className="flex justify-end gap-2">
        <DeleteConfirmSubmitButton onClick={remove} loading={loading} />
        <CancelButton onClick={close} disabled={loading} />
      </div>
    </BaseDialog>
  );
};
