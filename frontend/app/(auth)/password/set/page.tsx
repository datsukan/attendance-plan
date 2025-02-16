import type { Metadata } from 'next';

import { FormTitle } from '@/component/form/FormTitle';
import { Content } from './Content';

export const metadata: Metadata = {
  title: 'パスワード設定',
  description: 'TOU 受講計画管理のパスワード設定ページです。',
};

export default function PasswordSet() {
  return (
    <>
      <FormTitle label="パスワード設定" />
      <Content />
    </>
  );
}
