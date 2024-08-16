import { ReactNode } from 'react';

type Props = {
  children: ReactNode;
  onClick: () => void;
};

export const NavButton = ({ children, onClick }: Props) => {
  return (
    <button className="px-3 py-1.5 border rounded hover:bg-gray-50" onClick={() => onClick()}>
      {children}
    </button>
  );
};
