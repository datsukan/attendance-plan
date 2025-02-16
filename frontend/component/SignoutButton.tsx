'use client';

import { useRouter } from 'next/navigation';

import { removeAuthUser } from '@/storage/user';

export const SignoutButton = () => {
  const router = useRouter();

  const signout = () => {
    removeAuthUser();
    router.push('/signin');
  };

  return (
    <button className="rounded-md px-3 py-1 text-xs transition-colors hover:bg-gray-200" onClick={signout}>
      サインアウト
    </button>
  );
};
