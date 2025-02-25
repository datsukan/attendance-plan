'use client';

import { useState, useEffect } from 'react';
import { useSearchParams } from 'next/navigation';
import toast from 'react-hot-toast';

import { CompleteMessage } from './CompleteMessage';
import { Loading } from './Loading';

import { setEmail } from '@/backend-api/setEmail';
import { getUser } from '@/backend-api/getUser';
import { useStorage } from '@/provider/StorageProvider';

export const Content = () => {
  const searchParams = useSearchParams();
  const [isComplete, setIsComplete] = useState(false);
  const { user, saveUser } = useStorage();

  useEffect(() => {
    if (!user) return;
    if (isComplete) return;

    const idToken = searchParams.get('id_token') || '';
    const emailToken = searchParams.get('email_token') || '';

    if (!idToken || !emailToken) {
      toast.error('URLが不正です。');
      return;
    }

    (async () => {
      try {
        await setEmail(idToken, emailToken);
        setIsComplete(true);
        const newUser = await getUser();
        const newUserWithToken = { ...newUser, session_token: user.session_token };
        saveUser(newUserWithToken);
      } catch (e) {
        toast.error('メールアドレスの変更に失敗しました。');
        toast.error(String(e));
      }
    })();
  }, [searchParams, isComplete, user, saveUser]);

  return <div className="pt-16">{isComplete ? <CompleteMessage /> : <Loading />}</div>;
};
