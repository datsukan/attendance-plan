'use client';

import { useEffect, ReactNode } from 'react';
import { Dialog, DialogBackdrop, DialogPanel, DialogTitle } from '@headlessui/react';

type Props = {
  title: string;
  children: ReactNode;
  isOpen: boolean;
  disabled?: boolean;
  onClose: () => void;
};

export const BaseDialog = ({ children, title, isOpen, onClose, disabled = false }: Props) => {
  useEffect(() => {
    const downKey = (e: KeyboardEvent) => {
      if (e.key === 'Escape' && !disabled) {
        onClose();
      }
    };

    if (isOpen) {
      window.addEventListener('keydown', downKey);
    } else {
      window.removeEventListener('keydown', downKey);
    }

    return () => {
      window.removeEventListener('keydown', downKey);
    };
  }, [isOpen, disabled, onClose]);

  const close = () => {
    if (disabled) return;

    onClose();
  };

  return (
    <Dialog as="div" open={isOpen} onClose={close} className="relative z-50">
      <DialogBackdrop className="fixed inset-0 bg-black/30" />
      <div className="fixed inset-0 flex w-screen items-center justify-center p-4">
        <DialogPanel className="min-w-96 max-w-lg space-y-6 rounded-xl border bg-white p-6">
          <DialogTitle className="text-lg font-semibold">{title}</DialogTitle>
          {children}
        </DialogPanel>
      </div>
    </Dialog>
  );
};
