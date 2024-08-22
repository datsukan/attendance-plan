export type Schedule = {
  id: string;
  name: string;
  startDate: Date;
  endDate: Date;
  color: string;
  type: 'master' | 'custom';
};
