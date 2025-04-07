type ScheduleColorClassName = {
  [key: string]: string;
};

export const SCHEDULE_COLORS: ScheduleColorClassName[] = [
  { white: 'bg-white border border-gray-400' },
  { gray: 'bg-gray-500 border text-white' },
  { orange: 'bg-orange-500 text-white' },
  { blue: 'bg-blue-500 text-white' },
  { red: 'bg-red-500 text-white' },
  { green: 'bg-green-500 text-white' },
  { yellow: 'bg-yellow-500 text-white' },
  { outline_orange: 'bg-white border border-orange-500 text-orange-500' },
  { outline_blue: 'bg-white border border-blue-500 text-blue-500' },
  { outline_red: 'bg-white border border-red-500 text-red-500' },
  { outline_green: 'bg-white border border-green-500 text-green-500' },
  { outline_yellow: 'bg-white border border-yellow-500 text-yellow-500' },
];

export function getFirstColorKey(): string {
  return Object.keys(SCHEDULE_COLORS[0])[0];
}

export function getColorClassName(key: string): string {
  const color = SCHEDULE_COLORS.find((c) => c[key]);
  return color ? color[key] : '';
}

export function getColorKeys(): string[] {
  return SCHEDULE_COLORS.map((c) => Object.keys(c)[0]);
}
