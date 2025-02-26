'use client';

import { StorageProvider } from '@/provider/StorageProvider';
import { SubjectProvider } from '@/provider/SubjectProvider';

type Props = {
  children: React.ReactNode;
};

export const ProviderContainer = ({ children }: Props) => {
  return (
    <StorageProvider>
      <SubjectProvider>{children}</SubjectProvider>
    </StorageProvider>
  );
};
