'use client';

import { createContext, useContext, useState } from 'react';

type PopoverContextType = {
  isOpenPopover: boolean;
  openPopover: () => void;
  closePopover: () => void;
};

const createCtx = () => {
  const ctx = createContext<PopoverContextType | undefined>(undefined);
  const useCtx = () => {
    const c = useContext(ctx);
    if (!c) throw new Error('useCtx must be inside a Provider with a value');
    return c;
  };
  return [useCtx, ctx.Provider] as const;
};

const [useCtx, SetPopoverProvider] = createCtx();
export const usePopover = useCtx;

type Props = {
  children: React.ReactNode;
};

export const PopoverProvider = ({ children }: Props) => {
  const [isOpen, setIsOpen] = useState(false);

  const open = () => {
    setIsOpen(true);
  };

  const close = async () => {
    await new Promise((resolve) => setTimeout(resolve, 100));
    setIsOpen(false);
  };

  return <SetPopoverProvider value={{ isOpenPopover: isOpen, openPopover: open, closePopover: close }}>{children}</SetPopoverProvider>;
};
