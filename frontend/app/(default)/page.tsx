'use client';

import { Calender } from '@/component/calendar/Calendar';

import { useAuth } from '@/hook/useAuth';

export default function Home() {
  const [loadedAuth, isAuth] = useAuth();

  if (!loadedAuth || !isAuth) {
    return null;
  }

  return <Calender />;
}
