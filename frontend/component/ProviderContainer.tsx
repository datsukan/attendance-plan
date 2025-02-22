'use client';

import { StorageProvider } from '@/provider/StorageProvider';

type Props = {
  children: React.ReactNode;
};

export const ProviderContainer = ({ children }: Props) => {
  return <StorageProvider>{children}</StorageProvider>;
};
