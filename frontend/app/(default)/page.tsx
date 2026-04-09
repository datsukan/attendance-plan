'use client';

import { Calender } from '@/component/calendar/Calendar';

import { useAuth } from '@/hook/useAuth';
import { ScheduleProvider } from '@/provider/ScheduleProvider';
import { SelectionProvider } from '@/provider/SelectionContext';

export default function Home() {
  const [loadedAuth, isAuth] = useAuth();

  if (!loadedAuth || !isAuth) {
    return null;
  }

  return (
    <SelectionProvider>
      <ScheduleProvider>
        <Calender />
      </ScheduleProvider>
    </SelectionProvider>
  );
}
