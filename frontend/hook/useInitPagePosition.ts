import { useEffect } from 'react';

export const useInitPagePosition = (...deps: unknown[]) => {
  useEffect(() => {
    scrollTo(0, 0);
  }, [deps]);
};
