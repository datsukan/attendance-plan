'use client';

import { createContext, useContext, useRef } from 'react';
import toast from 'react-hot-toast';

import { UndoToast } from '@/component/UndoToast';

export type UndoCommand = {
  label: string;
  execute: () => Promise<void>;
};

type UndoContextType = {
  setUndoCommand: (command: UndoCommand) => void;
};

const UndoContext = createContext<UndoContextType | undefined>(undefined);

export const useUndo = (): UndoContextType => {
  const ctx = useContext(UndoContext);
  if (!ctx) throw new Error('useUndo must be used inside UndoProvider');
  return ctx;
};

type Props = {
  children: React.ReactNode;
};

export const UndoProvider = ({ children }: Props) => {
  const currentToastIdRef = useRef<string | null>(null);

  const setUndoCommand = (command: UndoCommand) => {
    if (currentToastIdRef.current) {
      toast.dismiss(currentToastIdRef.current);
    }

    const toastId = toast.custom(
      (t) => (
        <UndoToast
          toastId={t.id}
          label={command.label}
          onUndo={command.execute}
          visible={t.visible}
        />
      ),
      { duration: 15000 }
    );

    currentToastIdRef.current = toastId;
  };

  return (
    <UndoContext.Provider value={{ setUndoCommand }}>
      {children}
    </UndoContext.Provider>
  );
};
