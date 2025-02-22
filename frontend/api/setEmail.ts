import axios from 'axios';

import { loadAuthUser } from '@/storage/user';
import { newThrowResponseError } from './error';

export const setEmail = async (idToken: string, emailToken: string): Promise<void> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  const param = {
    id_token: idToken,
    email_token: emailToken,
  };

  try {
    await axios.post(`${process.env.NEXT_PUBLIC_API_BASE_URL}/email/set`, param, {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${user.session_token}`,
      },
    });

    return;
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
