import { ScheduleTypeMaster, ScheduleTypeCustom } from '@/const/schedule';

export type Schedule = {
  id: string;
  name: string;
  startDate: Date;
  endDate: Date;
  color: string;
  type: ScheduleType;
  order: number;
};

export type ScheduleType = typeof ScheduleTypeMaster | typeof ScheduleTypeCustom;
