import axios from 'axios';

import { newThrowResponseError } from './error';

export const resetPassword = async (email: string): Promise<void> => {
  const param = { email };

  try {
    await axios.post(`${process.env.NEXT_PUBLIC_API_BASE_URL}/password/reset`, param, {
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
