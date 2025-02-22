'use client';

import { LinkText } from '@/component/form/LinkText';

import { useStorage } from '@/provider/StorageProvider';

export const Email = () => {
  const { user } = useStorage();

  if (!user) return null;

  return (
    <div className="flex w-full flex-col gap-2">
      <label className="text-xs">メールアドレス</label>
      <div className="flex justify-between gap-2">
        <div title={user.email} className="truncate">
          {user.email}
        </div>
        <LinkText href="/email/reset">
          <span className="text-sm">変更</span>
        </LinkText>
      </div>
    </div>
  );
};
