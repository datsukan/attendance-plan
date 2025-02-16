export type ScheduleDateItem = {
  date: string;
  type: 'master' | 'custom';
  schedules: Schedule[];
};
