import axios from 'axios';
import { format } from 'date-fns';
import { Type } from '@/type';

import { loadAuthUser } from '@/storage/user';

export const updateBulkSchedules = async (schedules: Type.Schedule[]) => {
  const user = loadAuthUser();
  if (!user) {
    console.error('User not found');
    return;
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
  } catch (error) {
    console.error(error);
  }
};
