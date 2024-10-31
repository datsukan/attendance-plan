import { useEffect } from 'react';

export const usePagePosition = () => {
  useEffect(() => {
    scrollTo(0, 0);
  }, []);

  const init = () => {
    scrollTo(0, 0);
  };

  const execWhenPageBottom = (callback: () => void) => {
    const bodyHeight = document.body.scrollHeight;
    const windowHeight = window.innerHeight;
    const bottomPoint = bodyHeight - windowHeight - 200;

    const handleScroll = () => {
      const currentPos = window.scrollY;
      if (bottomPoint <= currentPos) {
        callback();
      }
    };

    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  };

  return {
    initPagePosition: init,
    execWhenPageBottom,
  };
};
