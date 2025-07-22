'use client';

import { PopoverProvider } from '@/provider/PopoverProvider';
import { UserProvider } from '@/provider/UserProvider';
import { SubjectProvider } from '@/provider/SubjectProvider';

type Props = {
  children: React.ReactNode;
};

export const ProviderContainer = ({ children }: Props) => {
  return (
    <PopoverProvider>
      <UserProvider>
        <SubjectProvider>{children}</SubjectProvider>
      </UserProvider>
    </PopoverProvider>
  );
};
