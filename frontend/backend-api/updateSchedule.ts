import axios from 'axios';
import { format } from 'date-fns';
import { Type } from '@/type';

import { loadAuthUser } from '@/storage/user';
import { newThrowResponseError } from './error';

export const updateSchedule = async (schedule: Type.Schedule): Promise<void> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  const param = {
    id: schedule.id,
    name: schedule.name,
    starts_at: format(schedule.startDate, 'yyyy-MM-dd HH:mm:ss'),
    ends_at: format(schedule.endDate, 'yyyy-MM-dd HH:mm:ss'),
    color: schedule.color,
    type: schedule.type,
    order: schedule.order,
  };

  try {
    await axios.put(`${process.env.NEXT_PUBLIC_API_BASE_URL}/schedules/${schedule.id}`, param, {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${user.session_token}`,
      },
    });
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
