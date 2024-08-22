type ScheduleColorClassName = {
  [key: string]: string;
};

export const SCHEDULE_COLORS: ScheduleColorClassName[] = [
  { white: 'bg-white border border-gray-400' },
  { orange: 'bg-orange-500 text-white' },
  { blue: 'bg-blue-500 text-white' },
  { red: 'bg-red-500 text-white' },
  { green: 'bg-green-500 text-white' },
  { yellow: 'bg-yellow-500 text-white' },
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
