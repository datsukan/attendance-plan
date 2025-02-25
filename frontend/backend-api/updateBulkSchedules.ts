import axios from 'axios';
import { format } from 'date-fns';
import { Type } from '@/type';

import { loadAuthUser } from '@/storage/user';
import { newThrowResponseError } from './error';

export const updateBulkSchedules = async (schedules: Type.Schedule[]): Promise<void> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  const params = schedules.map((schedule, i) => {
    return {
      id: schedule.id,
      name: schedule.name,
      starts_at: format(schedule.startDate, 'yyyy-MM-dd HH:mm:ss'),
      ends_at: format(schedule.endDate, 'yyyy-MM-dd HH:mm:ss'),
      color: schedule.color,
      type: schedule.type,
      order: i + 1,
    };
  });

  try {
    await axios.put(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/schedules/bulk`,
      { schedules: params },
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${user.session_token}`,
        },
      }
    );
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
