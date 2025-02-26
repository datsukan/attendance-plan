import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/react';
import { ChevronDownIcon } from '@heroicons/react/20/solid';

import { getColorClassName, getColorKeys } from '@/component/calendar/color-module';

type Props = {
  isSmall?: boolean;
  disabled?: boolean;
  value: string;
  onChange: (value: string) => void;
};

export const SelectColor = ({ isSmall = false, disabled = false, value, onChange }: Props) => {
  return (
    <Menu>
      <MenuButton className="flex items-center gap-1 rounded-lg p-2 hover:bg-gray-100 active:bg-gray-200" disabled={disabled}>
        <div className={`${isSmall ? 'size-4' : 'size-5'} rounded-full ${getColorClassName(value)}`} />
        <ChevronDownIcon className="size-4" />
      </MenuButton>

      <MenuItems transition anchor="bottom" className="grid grid-cols-2 gap-2 rounded border bg-white p-3 shadow-lg">
        {getColorKeys().map((key) => (
          <MenuItem key={key}>
            <button className={`size-6 rounded-full ${getColorClassName(key)}`} onClick={() => onChange(key)}></button>
          </MenuItem>
        ))}
      </MenuItems>
    </Menu>
  );
};
