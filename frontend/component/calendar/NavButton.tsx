import { ReactNode } from 'react';

type Props = {
  children: ReactNode;
  onClick: () => void;
};

export const NavButton = ({ children, onClick }: Props) => {
  return (
    <button className="px-3 py-0.5 border rounded hover:bg-gray-100 active:bg-gray-200" onClick={() => onClick()}>
      {children}
    </button>
  );
};
