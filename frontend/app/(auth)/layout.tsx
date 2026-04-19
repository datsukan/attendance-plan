'use client';

import Link from 'next/link';
import { QuestionMarkCircleIcon } from '@heroicons/react/24/outline';

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
            <div className="flex flex-col items-center gap-4">
              <div>直感的で手軽に受講スケジュールを管理するツール</div>
              <Link
                href="/guide"
                className="flex items-center gap-1 rounded-full bg-blue-50 px-4 py-1.5 text-sm text-blue-600 hover:bg-blue-100"
              >
                <QuestionMarkCircleIcon className="size-4" />
                このツールの特徴・機能
              </Link>
            </div>
            <div className="flex min-h-[25rem] w-full flex-col items-center gap-8 rounded-lg border p-4 shadow">{children}</div>
          </div>
        </main>
      )}
    </>
  );
}
