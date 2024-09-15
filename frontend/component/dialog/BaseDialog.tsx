import { ReactNode } from 'react';
import { Dialog, DialogBackdrop, DialogPanel, DialogTitle } from '@headlessui/react';

type Props = {
  children: ReactNode;
  isOpen: boolean;
  onClose: () => void;
  title: string;
};

export const BaseDialog = ({ children, isOpen, onClose, title }: Props) => {
  return (
    <Dialog as="div" open={isOpen} onClose={onClose} className="relative z-50">
      <DialogBackdrop className="fixed inset-0 bg-black/30" />
      <div className="fixed inset-0 flex w-screen items-center justify-center p-4">
        <DialogPanel className="max-w-lg min-w-96 space-y-6 border rounded-xl bg-white p-6">
          <DialogTitle className="text-lg font-semibold">{title}</DialogTitle>
          {children}
        </DialogPanel>
      </div>
    </Dialog>
  );
};
