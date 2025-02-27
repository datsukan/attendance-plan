'use client';

import { UserProvider } from '@/provider/UserProvider';
import { SubjectProvider } from '@/provider/SubjectProvider';

type Props = {
  children: React.ReactNode;
};

export const ProviderContainer = ({ children }: Props) => {
  return (
    <UserProvider>
      <SubjectProvider>{children}</SubjectProvider>
    </UserProvider>
  );
};
