'use client';

import { createContext, useContext, useState, useEffect } from 'react';

import { saveAuthUser, loadAuthUser, removeAuthUser } from '@/storage/user';
import type { AuthUser } from '@/storage/user';
import { getUser } from '@/backend-api/getUser';
import toast from 'react-hot-toast';

type UserContextType = {
  user: AuthUser | null;
  saveUser: (authUser: AuthUser) => void;
  removeUser: () => void;
};

const createCtx = () => {
  const ctx = createContext<UserContextType | undefined>(undefined);
  const useCtx = () => {
    const c = useContext(ctx);
    if (!c) throw new Error('useCtx must be inside a Provider with a value');
    return c;
  };
  return [useCtx, ctx.Provider] as const;
};

const [useCtx, SetUserProvider] = createCtx();
export const useUser = useCtx;

type Props = {
  children: React.ReactNode;
};

export const UserProvider = ({ children }: Props) => {
  const [user, setUser] = useState<AuthUser | null>(null);

  useEffect(() => {
    const au = loadAuthUser();
    if (!au) return;

    setUser(au);

    (async () => {
      try {
        const u = await getUser();
        const nu = {
          id: u.id,
          name: u.name,
          email: u.email,
          session_token: au.session_token,
        };
        setUser(nu);
      } catch (e) {
        toast.error('ユーザー情報の取得に失敗しました');
        toast.error(String(e));
        return;
      }
    })();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const save = (authUser: AuthUser) => {
    saveAuthUser(authUser);
    setUser(authUser);
  };

  const remove = () => {
    removeAuthUser();
    setUser(null);
  };

  return <SetUserProvider value={{ user, saveUser: save, removeUser: remove }}>{children}</SetUserProvider>;
};
