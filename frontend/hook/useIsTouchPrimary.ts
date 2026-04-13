'use client';

import { useState, useEffect } from 'react';

export const useIsTouchPrimary = (): boolean => {
  const [isTouchPrimary, setIsTouchPrimary] = useState(false);

  useEffect(() => {
    const mq = window.matchMedia('(pointer: coarse)');
    setIsTouchPrimary(mq.matches);

    const handler = (e: MediaQueryListEvent) => setIsTouchPrimary(e.matches);
    mq.addEventListener('change', handler);
    return () => mq.removeEventListener('change', handler);
  }, []);

  return isTouchPrimary;
};
