import { Suspense } from 'react';
import { FormTitle } from '@/component/form/FormTitle';
import { Content } from './Content';

export default function EmailSet() {
  return (
    <>
      <FormTitle label="メールアドレス設定" />
      <Suspense>
        <Content />
      </Suspense>
    </>
  );
}
