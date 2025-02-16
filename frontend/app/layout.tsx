import type { Metadata } from 'next';
import { M_PLUS_1 } from 'next/font/google';
import './globals.css';

const mPlus1 = M_PLUS_1({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: {
    template: '%s | TOU 受講計画管理',
    default: 'TOU 受講計画管理',
  },
  description: 'TOU向けの受講計画を管理するアプリケーションです。',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={mPlus1.className}>{children}</body>
    </html>
  );
}
