import axios from 'axios';
import { format } from 'date-fns';
import { Type } from '@/type';

import { loadAuthUser } from '@/storage/user';
import { newThrowResponseError } from './error';

export type CreateScheduleParam = Omit<Type.Schedule, 'id' | 'order'>;

export const createBulkSchedules = async (schedules: CreateScheduleParam[]): Promise<Type.Schedule[]> => {
  const user = loadAuthUser();
  if (!user) {
    throw new Error('User not found');
  }

  const params = schedules.map((schedule, i) => {
    return {
      name: schedule.name,
      starts_at: format(schedule.startDate, 'yyyy-MM-dd HH:mm:ss'),
      ends_at: format(schedule.endDate, 'yyyy-MM-dd HH:mm:ss'),
      color: schedule.color,
      type: schedule.type,
    };
  });

  try {
    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/schedules/bulk`,
      { schedules: params },
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${user.session_token}`,
        },
      }
    );

    if (!response.data?.schedules) {
      return [];
    }

    const s: Type.Schedule[] = response.data.schedules.map((data: any) => {
      return {
        id: data.id,
        name: data.name,
        startDate: new Date(data.starts_at),
        endDate: new Date(data.ends_at),
        color: data.color,
        type: data.type,
        order: data.order,
      };
    });

    return s;
  } catch (e) {
    newThrowResponseError(e);
    throw e;
  }
};
