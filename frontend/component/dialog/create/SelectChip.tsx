import { XMarkIcon } from '@heroicons/react/24/outline';

import { getColorClassName } from '@/component/calendar/color-module';

type Props = {
  label: string;
  color: string;
  onSelect: () => void;
  onRemove?: () => void;
};

export const SelectChip = ({ label, color, onSelect, onRemove }: Props) => {
  return (
    <div
      className={`flex max-w-full cursor-pointer items-center gap-1 rounded-full border pl-2 hover:bg-gray-100 active:bg-gray-200 ${
        onRemove ? 'pr-1.5' : 'pr-2'
      }`}
      onClick={onSelect}
    >
      <div className={`my-1 h-3 w-3 flex-shrink-0 rounded-full ${getColorClassName(color)}`} />
      <div className="my-1 truncate text-xs" title={label}>
        {label}
      </div>
      {onRemove && (
        <button
          type="button"
          className="rounded-full bg-gray-400 text-white hover:bg-gray-500 active:bg-gray-600"
          onClick={(e) => (e.stopPropagation(), onRemove())}
        >
          <XMarkIcon className="size-3.5" />
        </button>
      )}
    </div>
  );
};
