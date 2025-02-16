import { LinkText } from '@/component/form/LinkText';

export const CompleteMessage = () => {
  return (
    <div className="text-center leading-8">
      <p>パスワード設定が完了しました。</p>
      <p>
        <LinkText href="/signin">サインイン</LinkText>してください。
      </p>
    </div>
  );
};
