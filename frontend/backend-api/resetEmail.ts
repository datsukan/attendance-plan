import axios from 'axios';

import { loadAuthUser } from '@/storage/user';
import { newThrowResponseError } from './error';

export const resetEmail = async (email: string): Promise<void> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  const param = { email };

  try {
    await axios.post(`${process.env.NEXT_PUBLIC_API_BASE_URL}/users/${user.id}/email/reset`, param, {
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
