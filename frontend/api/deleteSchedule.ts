import axios from 'axios';

import { loadAuthUser } from '@/storage/user';

export const deleteSchedule = async (id: string) => {
  const user = loadAuthUser();
  if (!user) {
    console.error('User not found');
    return;
  }

  try {
    await axios.delete(`${process.env.NEXT_PUBLIC_API_BASE_URL}/schedules/${id}`, {
      headers: {
        Authorization: `Bearer ${user.session_token}`,
      },
    });
  } catch (error) {
    console.error(error);
  }
};
