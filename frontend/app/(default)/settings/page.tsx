'use client';

import { useAuth } from '@/hook/useAuth';

import { Info } from './Info';

export default function Settings() {
  const [loadedAuth, isAuth] = useAuth();

  if (!loadedAuth || !isAuth) {
    return null;
  }

  return (
    <div className="flex min-h-screen items-center justify-center p-2">
      <div className="flex min-h-[25rem] w-[30rem] flex-col items-center gap-8 rounded-lg border p-4 shadow">
        <h2 className="text-2xl">設定</h2>
        <div className="flex w-full flex-1">
          <Info />
        </div>
      </div>
    </div>
  );
}
