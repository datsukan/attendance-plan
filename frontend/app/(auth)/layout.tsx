'use client';

import { PageTitle } from '@/component/PageTitle';

import { useAuth } from '@/hook/useAuth';

export default function AuthLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const [loadedAuth, isAuth] = useAuth();

  return (
    <>
      {loadedAuth && !isAuth && (
        <main className="flex min-h-screen items-center justify-center p-2">
          <div className="flex w-[30rem] flex-col items-center gap-12">
            <PageTitle />
            <div>直感的で手軽に受講スケジュールを管理するツール</div>
            <div className="flex min-h-[25rem] w-full flex-col items-center gap-8 rounded-lg border p-4 shadow">{children}</div>
          </div>
        </main>
      )}
    </>
  );
}
