'use client';

import { useState, FormEventHandler } from 'react';
import { ArrowPathIcon } from '@heroicons/react/24/outline';
import toast from 'react-hot-toast';

import { updateUser } from '@/api/updateUser';
import { useStorage } from '@/provider/StorageProvider';

export const Name = () => {
  const [loading, setLoading] = useState(false);
  const { user, saveUser } = useStorage();

  if (!user) return null;

  const submit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    const form = new FormData(event.currentTarget);
    const name = form.get('name') || '';

    (async () => {
      setLoading(true);

      try {
        await updateUser(name.toString());
        saveUser({ ...user, name: name.toString() });
        toast.success('表示名を保存しました');
      } catch (e) {
        toast.error(String(e));
      }

      setLoading(false);
    })();
  };

  return (
    <form className="flex w-full items-end gap-2" onSubmit={submit}>
      <div className="flex w-full flex-col gap-2">
        <label className="text-xs" htmlFor="name">
          表示名
        </label>
        <div className="flex">
          <input id="name" name="name" type="name" defaultValue={user.name} disabled={loading} className="w-full rounded-s-lg border p-2" />
          <button
            type="submit"
            className="flex min-w-14 flex-shrink-0 items-center justify-center rounded-e-lg bg-blue-600 text-sm text-white hover:bg-blue-800 disabled:bg-blue-400"
            disabled={loading}
          >
            {loading ? <ArrowPathIcon className="size-5 animate-spin" /> : '保存'}
          </button>
        </div>
      </div>
    </form>
  );
};
