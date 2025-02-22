'use client';

import toast, { Toaster as ReactHotToaster, ToastBar } from 'react-hot-toast';
import { XMarkIcon } from '@heroicons/react/24/outline';

export const Toaster = () => {
  return (
    <ReactHotToaster position="bottom-right" toastOptions={{ duration: Infinity }}>
      {(t) => {
        return (
          <ToastBar toast={t} style={{ maxWidth: 'fit-content' }}>
            {({ icon, message }) => (
              <>
                {icon}
                <div className="mb-0.5">{message}</div>
                {t.type !== 'loading' && (
                  <button onClick={() => toast.dismiss(t.id)} className="rounded-full p-2 hover:bg-gray-200">
                    <XMarkIcon className="size-5" />
                  </button>
                )}
              </>
            )}
          </ToastBar>
        );
      }}
    </ReactHotToaster>
  );
};
