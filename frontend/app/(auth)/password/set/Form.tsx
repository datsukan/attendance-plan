'use client';

import { useState, FormEventHandler } from 'react';
import { useSearchParams } from 'next/navigation';

import { InputTextField } from '@/component/form/InputTextField';
import { SubmitButton } from '@/component/form/SubmitButton';

import { setPassword } from '@/api/setPassword';

type Props = {
  complete: () => void;
};

export const Form = ({ complete }: Props) => {
  const searchParams = useSearchParams();

  const [passwordErrorMessage, setPasswordErrorMessage] = useState('');
  const [rePasswordErrorMessage, setRePasswordErrorMessage] = useState('');
  const [loading, setLoading] = useState(false);

  const token = searchParams.get('token') || '';
  if (!token) {
    return (
      <span className="h-56 w-full font-bold text-red-500">
        URLが不正です。
        <br />
        パスワード設定のリクエストからやり直してください。
      </span>
    );
  }

  const submit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    const form = new FormData(event.currentTarget);
    const password = form.get('password') || '';
    const rePassword = form.get('re-password') || '';

    if (!password) {
      setPasswordErrorMessage('パスワードを入力してください');
      setRePasswordErrorMessage('');
      return;
    }

    if (!rePassword) {
      setPasswordErrorMessage('');
      setRePasswordErrorMessage('パスワードを再入力してください');
      return;
    }

    if (password !== rePassword) {
      setPasswordErrorMessage('');
      setRePasswordErrorMessage('パスワードが一致しません');
      return;
    }

    (async () => {
      setLoading(true);

      const errorMessage = await setPassword(token, password.toString());

      if (errorMessage) {
        setPasswordErrorMessage(errorMessage);
        setRePasswordErrorMessage('');
        setLoading(false);
        return;
      }

      setLoading(false);
      complete();
    })();
  };
  return (
    <form className="flex w-full flex-col gap-8" onSubmit={submit}>
      <div className="flex w-full flex-col gap-5">
        <InputTextField name="password" label="パスワード" type="password" errorMessage={passwordErrorMessage} />
        <InputTextField name="re-password" label="パスワードの再入力" type="password" errorMessage={rePasswordErrorMessage} />
      </div>
      <SubmitButton label="パスワードを設定する" loadingLabel="パスワードを設定中..." loading={loading} />
    </form>
  );
};
