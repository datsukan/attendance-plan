import type { Metadata } from 'next';
import { FormTitle } from '@/component/form/FormTitle';
import { Content } from './Content';

export const metadata: Metadata = {
  title: 'パスワードリセット',
  description: 'TOU 受講計画管理のパスワードリセットページです。',
};

export default function PasswordReset() {
  return (
    <>
      <FormTitle label="パスワードリセット" />
      <Content />
    </>
  );
}
