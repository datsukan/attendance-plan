import axios from 'axios';

import { loadAuthUser } from '@/storage/user';
import { newThrowResponseError } from './error';

type ResSubject = {
  id: string;
  user_id: string;
  name: string;
  color: string;
  created_at: string;
  updated_at: string;
};

export const createSubject = async (name: string, color: string): Promise<ResSubject> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  const param = {
    name,
    color,
  };

  try {
    const response = await axios.post(`${process.env.NEXT_PUBLIC_API_BASE_URL}/subjects`, param, {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${user.session_token}`,
      },
    });

    return response.data;
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
