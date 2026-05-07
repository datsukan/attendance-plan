import axios from 'axios';

import { loadAuthUser } from '@/storage/user';
import { newThrowResponseError } from './error';

export type UserUsageSubject = {
  id: string;
  name: string;
  color: string;
  created_at: string;
  updated_at: string;
};

export type UserUsage = {
  id: string;
  email: string;
  name: string;
  registered_at: string;
  last_used_at: string;
  subjects: UserUsageSubject[];
};

export const fetchUserUsages = async (): Promise<UserUsage[]> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  try {
    const response = await axios.get(`${process.env.NEXT_PUBLIC_API_BASE_URL}/user-usages`, {
      headers: {
        Authorization: `Bearer ${user.session_token}`,
      },
    });

    if (!response.data || !response.data.users) {
      return [];
    }

    return response.data.users;
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
