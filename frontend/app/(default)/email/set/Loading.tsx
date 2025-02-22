import { ArrowPathIcon } from '@heroicons/react/24/outline';

export const Loading = () => {
  return (
    <div className="flex items-center gap-2">
      <ArrowPathIcon className="size-8 animate-spin" />
      メールアドレスの変更を認証中...
    </div>
  );
};
