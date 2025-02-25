import type { Metadata } from 'next';
import { FormTitle } from '@/component/form/FormTitle';
import { Form } from './Form';
import { Note } from './Note';

export const metadata: Metadata = {
  title: 'サインイン',
  description: 'TOU 受講スケジュール管理のサインインページです。',
};

export default function SignIn() {
  return (
    <>
      <FormTitle label="サインイン" />
      <Form />
      <Note />
    </>
  );
}
