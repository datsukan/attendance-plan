'use client';

import { useState, useRef, KeyboardEvent } from 'react';
import { PlusIcon, ArrowPathIcon } from '@heroicons/react/24/outline';

import { SelectColor } from '../SelectColor';
import { getFirstColorKey } from '@/component/calendar/color-module';

type Props = {
  loading?: boolean;
  submit: (name: string, color: string) => Promise<void>;
};

export const AddChip = ({ loading = false, submit }: Props) => {
  const [name, setName] = useState('');
  const [colorKey, setColorKey] = useState(getFirstColorKey());
  const inputRef = useRef<HTMLInputElement>(null);

  const keyDown = (e: KeyboardEvent<HTMLInputElement>) => {
    if (loading) return;
    if (e.nativeEvent.isComposing || e.key !== 'Enter') return;
    add();
  };

  const add = async () => {
    if (loading) return;
    if (!name) return;

    await submit(name, colorKey);
    setName('');
    setColorKey(getFirstColorKey());

    setTimeout(() => {
      inputRef.current?.focus();
    }, 100);
  };

  return (
    <div className="flex w-full items-center gap-1 rounded-full border pl-4 pr-1 active:ring-1">
      <input
        ref={inputRef}
        type="text"
        value={name}
        autoComplete="schedule-label"
        className="w-full border-b text-sm focus-visible:border-blue-500 focus-visible:outline-none"
        placeholder="テンプレート用のスケジュール名"
        onChange={(e) => setName(e.target.value)}
        onKeyDown={keyDown}
        disabled={loading}
      />
      <SelectColor isSmall value={colorKey} onChange={setColorKey} disabled={loading} />
      <button
        type="button"
        title="追加"
        className="rounded-full bg-blue-600 p-0.5 hover:bg-blue-500 active:bg-blue-800 disabled:bg-blue-400"
        onClick={add}
        disabled={loading}
      >
        {loading ? <ArrowPathIcon className="size-5 animate-spin text-white" /> : <PlusIcon className="size-5 text-white" />}
      </button>
    </div>
  );
};
