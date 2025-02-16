import axios from 'axios';
import { format } from 'date-fns';
import { Type } from '@/type';

import { loadAuthUser } from '@/storage/user';

export type CreateScheduleParam = Omit<Type.Schedule, 'id' | 'order'>;

export const createSchedule = async (schedule: CreateScheduleParam): Promise<Type.Schedule | null> => {
  const user = loadAuthUser();
  if (!user) {
    console.error('User not found');
    return null;
  }

  const param = {
    name: schedule.name,
    starts_at: format(schedule.startDate, 'yyyy-MM-dd HH:mm:ss'),
    ends_at: format(schedule.endDate, 'yyyy-MM-dd HH:mm:ss'),
    color: schedule.color,
    type: schedule.type,
  };

  try {
    const response = await axios.post(`${process.env.NEXT_PUBLIC_API_BASE_URL}/schedules`, param, {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${user.session_token}`,
      },
    });

    const s: Type.Schedule = {
      id: response.data.id,
      name: response.data.name,
      startDate: new Date(response.data.starts_at),
      endDate: new Date(response.data.ends_at),
      color: response.data.color,
      type: response.data.type,
      order: response.data.order,
    };

    return s;
  } catch (error) {
    console.error(error);
  }

  return null;
};
