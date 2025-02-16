'use client';

import { useEffect, useState } from 'react';
import { useRouter, usePathname } from 'next/navigation';

import { loadAuthUser } from '@/storage/user';

export const useAuth = () => {
  const router = useRouter();
  const pathname = usePathname();
  const [loaded, setLoaded] = useState(false);
  const [isAuth, setIsAuth] = useState(false);

  useEffect(() => {
    const user = loadAuthUser();
    setLoaded(true);

    if (user) {
      setIsAuth(true);
      router.push('/');
      return;
    }

    setIsAuth(false);

    switch (pathname) {
      case '/signup':
      case '/password/set':
      case '/password/reset':
        break;
      default:
        router.push('/signin');
        break;
    }
  }, [router, pathname]);

  return [loaded, isAuth];
};
