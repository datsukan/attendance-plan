'use client';

import { PopoverProvider } from '@/provider/PopoverProvider';
import { UserProvider } from '@/provider/UserProvider';
import { SubjectProvider } from '@/provider/SubjectProvider';
import { LoadingProvider } from '@/provider/LoadingProvider';
import { LoadingBar } from '@/component/LoadingBar';

type Props = {
  children: React.ReactNode;
};

export const ProviderContainer = ({ children }: Props) => {
  return (
    <LoadingProvider>
      <LoadingBar />
      <PopoverProvider>
        <UserProvider>
          <SubjectProvider>{children}</SubjectProvider>
        </UserProvider>
      </PopoverProvider>
    </LoadingProvider>
  );
};
