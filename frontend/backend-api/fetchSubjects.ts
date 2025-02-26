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

export const fetchSubjects = async (): Promise<ResSubject[]> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  try {
    const response = await axios.get(`${process.env.NEXT_PUBLIC_API_BASE_URL}/users/${user.id}/subjects`, {
      headers: {
        Authorization: `Bearer ${user.session_token}`,
      },
    });

    if (!response.data || !response.data.subjects) {
      return [];
    }

    return response.data.subjects;
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
