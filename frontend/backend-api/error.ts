import axios from 'axios';

import { removeAuthUser } from '@/storage/user';

export const UnknownErrorMessage = 'エラーが発生しました';

export const newThrowResponseError = (e: unknown): void => {
  if (!axios.isAxiosError(e)) {
    console.error(e);
    throw new Error(UnknownErrorMessage);
  }

  if (e.response?.status === 401) {
    removeAuthUser();
  }

  const body = e.response?.data;
  if (!body) {
    console.error('response body is empty');
    throw new Error(UnknownErrorMessage);
  }

  throw new Error(body.message || UnknownErrorMessage);
};
