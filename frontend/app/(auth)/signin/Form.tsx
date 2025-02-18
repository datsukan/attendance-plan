'use client';

import { useState, FormEventHandler } from 'react';
import { useRouter } from 'next/navigation';

import { InputTextField } from '@/component/form/InputTextField';
import { SubmitButton } from '@/component/form/SubmitButton';

import { signin } from '@/api/signin';
import { saveAuthUser } from '@/storage/user';

export const Form = () => {
  const router = useRouter();

  const [emailErrorMessage, setEmailErrorMessage] = useState('');
  const [passwordErrorMessage, setPasswordErrorMessage] = useState('');
  const [loading, setLoading] = useState(false);

  const submit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    const form = new FormData(event.currentTarget);
    const email = form.get('email') || '';
    const password = form.get('password') || '';
    let hasError = false;

    if (!email) {
      setEmailErrorMessage('メールアドレスを入力してください');
      hasError = true;
    }

    if (!password) {
      setPasswordErrorMessage('パスワードを入力してください');
      hasError = true;
    }

    if (hasError) {
      return;
    }

    (async () => {
      setLoading(true);

      try {
        const user = await signin(email.toString(), password.toString());
        saveAuthUser({
          id: user.id,
          email: user.email,
          name: user.name,
          session_token: user.sessionToken,
        });
      } catch (e) {
        setLoading(false);

        if (e instanceof Error) {
          setEmailErrorMessage(e.message);
          setPasswordErrorMessage('');
          return;
        }

        alert(String(e));
        return;
      }

      setLoading(false);
      setEmailErrorMessage('');
      setPasswordErrorMessage('');
      router.push('/');
    })();
  };
  return (
    <form className="flex w-full flex-col gap-8" onSubmit={submit}>
      <div className="flex w-full flex-col gap-5">
        <InputTextField name="email" label="メールアドレス" type="email" errorMessage={emailErrorMessage} />
        <InputTextField name="password" label="パスワード" type="password" errorMessage={passwordErrorMessage} />
      </div>
      <SubmitButton label="サインインする" loadingLabel="サインイン中..." loading={loading} />
    </form>
  );
};
