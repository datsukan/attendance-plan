import axios from 'axios';

import { loadAuthUser } from '@/storage/user';
import { newThrowResponseError } from './error';

export const deleteSchedule = async (id: string): Promise<void> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  try {
    await axios.delete(`${process.env.NEXT_PUBLIC_API_BASE_URL}/schedules/${id}`, {
      headers: {
        Authorization: `Bearer ${user.session_token}`,
      },
    });
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
