'use client';

import { useState, FormEventHandler } from 'react';
import toast from 'react-hot-toast';

import { InputTextField } from '@/component/form/InputTextField';
import { SubmitButton } from '@/component/form/SubmitButton';

import { resetEmail } from '@/backend-api/resetEmail';

type Props = {
  complete: () => void;
};

export const Form = ({ complete }: Props) => {
  const [errorMessage, setErrorMessage] = useState('');
  const [loading, setLoading] = useState(false);

  const submit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    const form = new FormData(event.currentTarget);
    const email = form.get('email') || '';

    if (!email) {
      setErrorMessage('メールアドレスを入力してください');
      return;
    }

    (async () => {
      setLoading(true);

      try {
        await resetEmail(email.toString());
      } catch (e) {
        setLoading(false);

        if (e instanceof Error) {
          setErrorMessage(e.message);
          return;
        }

        toast.error(String(e));
        return;
      }

      setLoading(false);
      complete();
    })();
  };

  return (
    <form className="flex w-full flex-col gap-8 pt-8" onSubmit={submit}>
      <InputTextField name="email" label="新しいメールアドレス" type="email" autocomplete="current-password" errorMessage={errorMessage} />
      <SubmitButton label="変更する" loadingLabel="変更中..." loading={loading} />
    </form>
  );
};
