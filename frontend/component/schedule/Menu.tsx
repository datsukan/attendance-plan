import { TrashIcon, PencilIcon } from '@heroicons/react/24/outline';

import { getColorClassName, getColorKeys } from '@/component/calendar/color-module';

type Props = {
  onSelectColor: (color: string) => void;
  openRemoveConfirmDialog: () => void;
  openEditDialog: () => void;
};

export const Menu = ({ onSelectColor, openRemoveConfirmDialog, openEditDialog }: Props) => {
  return (
    <div className="overflow-hidden rounded-lg bg-white shadow-lg">
      <button className="flex items-center gap-2 px-4 py-2 hover:bg-gray-100 active:bg-gray-200" onClick={() => openEditDialog()}>
        <PencilIcon className="size-5 text-gray-600" />
        <span>編集</span>
      </button>
      <button className="flex items-center gap-2 px-4 py-2 hover:bg-gray-100 active:bg-gray-200" onClick={openRemoveConfirmDialog}>
        <TrashIcon className="size-5 text-gray-600" />
        <span>削除</span>
      </button>
      <div className="grid grid-cols-2 gap-2 p-2">
        {getColorKeys().map((color) => (
          <button key={color} className={`mx-auto size-6 rounded-full ${getColorClassName(color)}`} onClick={() => onSelectColor(color)} />
        ))}
      </div>
    </div>
  );
};
