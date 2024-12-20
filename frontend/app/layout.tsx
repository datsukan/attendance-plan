import type { Metadata } from "next";
import { M_PLUS_1 } from "next/font/google";
import "./globals.css";

const mPlus1 = M_PLUS_1({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
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
