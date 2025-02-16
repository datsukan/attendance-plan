'use client';

import { PageTitle } from '@/component/PageTitle';
import { SignoutButton } from '@/component/SignoutButton';
import { Calender } from '@/component/calendar/Calendar';

import { useAuth } from '@/hook/useAuth';

export default function Home() {
  const [loadedAuth, isAuth] = useAuth();

  if (!loadedAuth || !isAuth) {
    return null;
  }

  return (
    <>
      <main className="container mx-auto px-4 py-8">
        <div className="flex content-center justify-between">
          <PageTitle />
          <SignoutButton />
        </div>
        <div className="mt-4">
          <Calender />
        </div>
      </main>
    </>
  );
}
