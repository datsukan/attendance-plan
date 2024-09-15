import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/react';
import { ChevronDownIcon } from '@heroicons/react/20/solid';

import { getColorClassName, getColorKeys } from '@/component/calendar/color-module';

type Props = {
  value: string;
  onChange: (value: string) => void;
};

export const SelectColor = ({ value, onChange }: Props) => {
  return (
    <Menu>
      <MenuButton className="rounded-lg p-2 flex gap-1 items-center hover:bg-gray-100 active:bg-gray-200">
        <div className={`size-5 rounded-full ${getColorClassName(value)}`} />
        <ChevronDownIcon className="size-4" />
      </MenuButton>

      <MenuItems transition anchor="bottom" className="border shadow-lg rounded p-3 bg-white grid grid-cols-2 gap-2">
        {getColorKeys().map((key) => (
          <MenuItem key={key}>
            <button className={`size-6 rounded-full ${getColorClassName(key)}`} onClick={() => onChange(key)}></button>
          </MenuItem>
        ))}
      </MenuItems>
    </Menu>
  );
};
