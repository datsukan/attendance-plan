'use client';

import { useSelection } from '@/provider/SelectionContext';
import { useIsTouchPrimary } from '@/hook/useIsTouchPrimary';

type Props = {
  onDelete: () => void;
};

export const SelectionModeBar = ({ onDelete }: Props) => {
  const { selectedIds, exitSelectionMode } = useSelection();
  const isTouchPrimary = useIsTouchPrimary();
  const count = selectedIds.size;

  return (
    <div className="flex items-center justify-between bg-blue-50 px-4 py-2 text-sm">
      <div>
        <span className="font-medium text-blue-700">{count}件選択中</span>
        <p className="mt-0.5 text-xs text-blue-400">
          {isTouchPrimary ? '長押しで選択モード開始 / タップで追加・解除' : 'Ctrl+クリック で追加選択 / Shift+クリック で範囲選択'}
        </p>
      </div>
      <div className="flex gap-2">
        <button
          onClick={exitSelectionMode}
          className="rounded px-3 py-1 text-blue-700 hover:bg-blue-100"
        >
          キャンセル
        </button>
        <button
          onClick={onDelete}
          disabled={count === 0}
          className="rounded bg-red-500 px-3 py-1 text-white hover:bg-red-600 disabled:opacity-40"
        >
          削除
        </button>
      </div>
    </div>
  );
};
