import axios from 'axios';

import { newThrowResponseError } from './error';

export const signup = async (email: string): Promise<void> => {
  const param = { email };

  try {
    await axios.post(`${process.env.NEXT_PUBLIC_API_BASE_URL}/signup`, param, {
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
