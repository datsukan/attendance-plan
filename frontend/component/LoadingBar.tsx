'use client';

import { useLoadingBar } from '@/provider/LoadingProvider';

export const LoadingBar = () => {
  const { isLoading } = useLoadingBar();

  if (!isLoading) return null;

  return (
    <div className="fixed left-0 top-0 z-50 h-1 w-full overflow-hidden bg-blue-100">
      <div className="animate-loading-bar h-full w-1/2 rounded-full bg-blue-500" />
    </div>
  );
};
