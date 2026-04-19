import axios from 'axios';
import toast from 'react-hot-toast';

import { removeAuthUser } from '@/storage/user';

export const UnknownErrorMessage = 'エラーが発生しました';

export class SessionExpiredError extends Error {
  constructor() {
    super('セッションが切れました');
    this.name = 'SessionExpiredError';
  }
}

let navigate: ((path: string) => void) | null = null;

export const registerNavigate = (fn: (path: string) => void) => {
  navigate = fn;
};

export const newThrowResponseError = (e: unknown): void => {
  if (!axios.isAxiosError(e)) {
    console.error(e);
    throw new Error(UnknownErrorMessage);
  }

  if (e.response?.status === 401) {
    removeAuthUser();
    toast.error('セッションが切れました。再ログインしてください。', { id: 'session-expired' });
    navigate?.('/signin');
    throw new SessionExpiredError();
  }

  const body = e.response?.data;
  if (!body) {
    console.error('response body is empty');
    throw new Error(UnknownErrorMessage);
  }

  throw new Error(body.message || UnknownErrorMessage);
};
