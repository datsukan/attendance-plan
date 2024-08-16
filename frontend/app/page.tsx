import { Calender } from '@/component/calendar/Calendar';

export default function Home() {
  return (
    <main className="container mx-auto px-4 py-8">
      <h1 className="text-2xl font-semibold">受講計画</h1>
      <div className="mt-4">
        <Calender />
      </div>
    </main>
  );
}
