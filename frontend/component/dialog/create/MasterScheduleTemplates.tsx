import { getMasterScheduleTemplates } from '@/component/schedule/schedule-module';
import { getColorClassName } from '@/component/calendar/color-module';

type Props = {
  onSelect: (name: string, color: string) => void;
};

export const MasterScheduleTemplates = ({ onSelect }: Props) => {
  const templates = getMasterScheduleTemplates();

  return (
    <div className="flex gap-1 flex-wrap">
      {templates.map((item, index) => (
        <button
          key={index}
          className="border rounded-full px-3 py-1 flex gap-1 items-center hover:bg-gray-100 active:bg-gray-200"
          onClick={() => onSelect(item.name, item.color)}
        >
          <div className={`w-3 h-3 rounded-full ${getColorClassName(item.color)}`} />
          <div className="text-xs">{item.name}</div>
        </button>
      ))}
    </div>
  );
};
