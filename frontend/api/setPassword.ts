import axios from 'axios';

import { newThrowResponseError } from './error';

export const setPassword = async (token: string, password: string): Promise<void> => {
  const param = { token, password };

  try {
    await axios.post(`${process.env.NEXT_PUBLIC_API_BASE_URL}/password/set`, param, {
      headers: {
        'Content-Type': 'application/json',
      },
    });

    return;
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
