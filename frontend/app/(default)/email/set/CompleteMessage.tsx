import { LinkText } from '@/component/form/LinkText';

export const CompleteMessage = () => {
  return (
    <div className="text-center leading-8">
      <p>メールアドレスの変更が完了しました。</p>
      <p>
        <LinkText href="/settings">設定</LinkText>から確認できます。
      </p>
    </div>
  );
};
