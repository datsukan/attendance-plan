export const DayOfWeeks = () => {
  const items = ['月', '火', '水', '木', '金', '土', '日'];
  return (
    <div className="border-r border-b grid grid-cols-7">
      {items.map((item) => {
        return (
          <div
            key={item}
            className={`min-h-10 text-center p-1 border-t border-l flex justify-center items-center ${isHoliday(item) ? 'bg-gray-50' : ''}`}
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
