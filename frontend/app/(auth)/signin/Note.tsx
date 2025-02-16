import { LinkText } from '@/component/form/LinkText';

export const Note = () => {
  return (
    <div className="w-full text-xs leading-6 text-gray-500">
      <p>
        未登録の場合は
        <LinkText href="/signup">サインアップ</LinkText>
        してください。
      </p>
      <p>
        パスワードを忘れた方は<LinkText href="/password/reset">パスワードリセット</LinkText>してください。
      </p>
    </div>
  );
};
