import axios from 'axios';
import { Type } from '@/type';

import { loadAuthUser } from '@/storage/user';

export type FetchScheduleResult = {
  masterSchedules: Type.ScheduleDateItem[];
  customSchedules: Type.ScheduleDateItem[];
};

type ResSchedule = {
  id: string;
  name: string;
  starts_at: string;
  ends_at: string;
  color: string;
  type: string;
  order: number;
};

type ResScheduleDateItem = {
  date: string;
  type: string;
  schedules: ResSchedule[];
};

type ResScheduleDateItemList = ResScheduleDateItem[];

export const fetchSchedules = async (): Promise<FetchScheduleResult> => {
  const user = loadAuthUser();
  if (!user) {
    console.error('User not found');
    return { masterSchedules: [], customSchedules: [] };
  }

  try {
    const response = await axios.get(`${process.env.NEXT_PUBLIC_API_BASE_URL}/users/${user.id}/schedules`, {
      headers: {
        Authorization: `Bearer ${user.session_token}`,
      },
    });

    let result = { masterSchedules: [], customSchedules: [] } as FetchScheduleResult;

    if (response.data.master_schedules) {
      result.masterSchedules = toResultSchedules(response.data.master_schedules);
    }

    if (response.data.custom_schedules) {
      result.customSchedules = toResultSchedules(response.data.custom_schedules);
    }

    return result;
  } catch (error) {
    console.error(error);
  }

  return { masterSchedules: [], customSchedules: [] };
};

const toResultSchedules = (schedules: ResScheduleDateItemList): Type.ScheduleDateItem[] => {
  let result: Type.ScheduleDateItem[] = [];
  for (const s of schedules) {
    result.push({
      date: s.date,
      type: s.type as 'master' | 'custom',
      schedules: s.schedules.map((ss) => ({
        id: ss.id,
        name: ss.name,
        startDate: new Date(ss.starts_at),
        endDate: new Date(ss.ends_at),
        color: ss.color,
        type: ss.type,
        order: ss.order,
      })),
    });
  }

  return result;
};
