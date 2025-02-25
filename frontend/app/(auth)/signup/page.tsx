import type { Metadata } from 'next';
import { FormTitle } from '@/component/form/FormTitle';
import { Content } from './Content';

export const metadata: Metadata = {
  title: 'サインアップ',
  description: 'TOU 受講スケジュール管理のサインアップページです。',
};

export default function SignUp() {
  return (
    <>
      <FormTitle label="サインアップ" />
      <Content />
    </>
  );
}
