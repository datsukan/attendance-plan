'use client';

import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/react';
import {
  ChevronDownIcon,
  UserCircleIcon,
  Cog6ToothIcon,
  ArrowRightStartOnRectangleIcon,
  EllipsisHorizontalIcon,
} from '@heroicons/react/20/solid';

import { useUser } from '@/provider/UserProvider';

export const AccountDropDown = () => {
  const router = useRouter();
  const { user, removeUser } = useUser();

  if (!user) return null;

  const signout = () => {
    removeUser();
    router.push('/signin');
  };

  return (
    <Menu as="div" className="relative min-w-fit">
      <MenuButton className="rounded py-1 ring-1 ring-gray-200 hover:bg-gray-100">
        <div className="hidden items-center space-x-2 pl-3 pr-1 sm:flex">
          <UserCircleIcon className="size-5" />
          <span className="mb-0.5 text-sm">{!user.name ? user.email : user.name}</span>
          <ChevronDownIcon className="size-5" />
        </div>
        <div className="px-2 sm:hidden">
          <EllipsisHorizontalIcon className="size-6" />
        </div>
      </MenuButton>

      <MenuItems
        as="div"
        className="absolute right-0 z-50 mt-2 w-48 origin-top-right divide-y divide-gray-100 rounded-md bg-white text-sm shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
        transition
      >
        <Link href="/settings">
          <MenuItem as="div" className="flex items-center gap-2 px-4 py-2 data-[focus]:bg-gray-100">
            <Cog6ToothIcon className="size-4" />
            <span className="mb-0.5">設定</span>
          </MenuItem>
        </Link>
        <MenuItem as="button" className="flex w-full items-center gap-2 px-4 py-2 text-left data-[focus]:bg-gray-100" onClick={signout}>
          <ArrowRightStartOnRectangleIcon className="size-4" />
          <span className="mb-0.5">サインアウト</span>
        </MenuItem>
      </MenuItems>
    </Menu>
  );
};
