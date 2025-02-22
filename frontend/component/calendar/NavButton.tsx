import { ReactNode } from 'react';

type Props = {
  children: ReactNode;
  onClick: () => void;
};

export const NavButton = ({ children, onClick }: Props) => {
  return (
    <button className="rounded border px-5 py-1.5 hover:bg-gray-100 active:bg-gray-200" onClick={() => onClick()}>
      {children}
    </button>
  );
};
