import type { Metadata } from 'next';

import { PageTitle } from '@/component/PageTitle';
import { AccountDropDown } from '@/component/AccountDropDown';

export const metadata: Metadata = {
  title: {
    template: '%s | TOU 受講スケジュール管理',
    default: 'TOU 受講スケジュール管理',
  },
  description: 'TOU向けの受講スケジュールを管理するアプリケーションです。',
  openGraph: {
    title: 'TOU 受講スケジュール管理',
    description: 'TOU向けの受講スケジュールを管理するアプリケーションです。',
  },
  twitter: {
    title: 'TOU 受講スケジュール管理',
    description: 'TOU向けの受講スケジュールを管理するアプリケーションです。',
    card: 'summary_large_image',
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body>
        <main className="container mx-auto px-4 py-8">
          <div className="flex flex-wrap content-center justify-between gap-2">
            <PageTitle />
            <AccountDropDown />
          </div>
          <div className="mt-4">{children}</div>
        </main>
      </body>
    </html>
  );
}
