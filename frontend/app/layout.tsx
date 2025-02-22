import { M_PLUS_1 } from 'next/font/google';
import './globals.css';
import { Toaster } from '@/component/Toaster';
import { ProviderContainer } from '@/component/ProviderContainer';

const mPlus1 = M_PLUS_1({ subsets: ['latin'] });

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={mPlus1.className}>
        <Toaster />
        <ProviderContainer>{children}</ProviderContainer>
      </body>
    </html>
  );
}
