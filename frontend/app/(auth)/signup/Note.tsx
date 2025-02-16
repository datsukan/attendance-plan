import { LinkText } from '@/component/form/LinkText';

export const Note = () => {
  return (
    <div className="w-full text-xs leading-6 text-gray-500">
      <p>登録したメールアドレスは本サイト内での制御のみに使用されます。</p>
      <p>パスワードはハッシュ化されて保存され、他者には知られません。</p>
      <p>
        登録を実行すると入力したメールアドレス宛てにパスワード設定ページへのリンクが送られるので、そちらからパスワードを設定して登録を完了してください。
      </p>
      <p>
        既に登録済みの場合は
        <LinkText href="/signin">サインイン</LinkText>
        してください。
      </p>
    </div>
  );
};
