export const DayOfWeeks = () => {
  const items = ['月', '火', '水', '木', '金', '土', '日'];
  return (
    <div className="grid grid-cols-7 border-b border-r">
      {items.map((item) => {
        return (
          <div
            key={item}
            className={`flex min-h-10 items-center justify-center border-l border-t p-1 text-center ${isHoliday(item) ? 'bg-gray-50' : ''}`}
          >
            <span className="text-sm">{item}</span>
          </div>
        );
      })}
    </div>
  );
};

function isHoliday(item: string) {
  return item === '土' || item === '日';
}
