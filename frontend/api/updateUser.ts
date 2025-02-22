import axios from 'axios';

import { loadAuthUser } from '@/storage/user';
import { newThrowResponseError } from './error';

export const updateUser = async (name: string): Promise<void> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  const param = { name };

  try {
    const response = await axios.put(`${process.env.NEXT_PUBLIC_API_BASE_URL}/users/${user.id}`, param, {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${user.session_token}`,
      },
    });

    if (!response.data) {
      throw new Error('Response data not found');
    }
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
