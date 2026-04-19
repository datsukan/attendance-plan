'use client';

import { createContext, useCallback, useContext, useRef, useState } from 'react';

type LoadingContextType = {
  isLoading: boolean;
  startLoading: () => void;
  stopLoading: () => void;
};

const LoadingContext = createContext<LoadingContextType | undefined>(undefined);

export const useLoadingBar = () => {
  const c = useContext(LoadingContext);
  if (!c) throw new Error('useLoadingBar must be inside LoadingProvider');
  return c;
};

type Props = {
  children: React.ReactNode;
};

export const LoadingProvider = ({ children }: Props) => {
  const [isLoading, setIsLoading] = useState(false);
  const countRef = useRef(0);

  const startLoading = useCallback(() => {
    countRef.current += 1;
    setIsLoading(true);
  }, []);

  const stopLoading = useCallback(() => {
    countRef.current = Math.max(0, countRef.current - 1);
    if (countRef.current === 0) {
      setIsLoading(false);
    }
  }, []);

  return <LoadingContext.Provider value={{ isLoading, startLoading, stopLoading }}>{children}</LoadingContext.Provider>;
};
