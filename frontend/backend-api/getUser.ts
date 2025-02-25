import axios from 'axios';

import { loadAuthUser } from '@/storage/user';
import { newThrowResponseError } from './error';

type result = {
  id: string;
  email: string;
  name: string;
  createdAt: string;
  updatedAt: string;
};

export const getUser = async (): Promise<result> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  try {
    const response = await axios.get(`${process.env.NEXT_PUBLIC_API_BASE_URL}/users/${user.id}`, {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${user.session_token}`,
      },
    });

    const result: result = {
      id: response.data.id,
      email: response.data.email,
      name: response.data.name,
      createdAt: response.data.created_at,
      updatedAt: response.data.updated_at,
    };

    return result;
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
