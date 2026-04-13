'use client';

import { createContext, useContext, useState, ReactNode } from 'react';

type ScheduleRef = { id: string; date: string };

type SelectionContextType = {
  selectedIds: Set<string>;
  toggleSelect: (id: string) => void;
  rangeSelect: (id: string) => void;
  clearSelection: () => void;
  isSelected: (id: string) => boolean;
  setAllSchedules: (schedules: ScheduleRef[]) => void;
  isSelectionMode: boolean;
  enterSelectionMode: (firstId: string) => void;
  exitSelectionMode: () => void;
};

const SelectionContext = createContext<SelectionContextType | null>(null);

export const SelectionProvider = ({ children }: { children: ReactNode }) => {
  const [selectedIds, setSelectedIds] = useState<Set<string>>(new Set());
  const [anchorId, setAnchorId] = useState<string | null>(null);
  const [selectionDate, setSelectionDate] = useState<string | null>(null);
  const [allSchedules, setAllSchedules] = useState<ScheduleRef[]>([]);
  const [isSelectionMode, setIsSelectionMode] = useState(false);

  const getDate = (id: string) => allSchedules.find((s) => s.id === id)?.date ?? null;

  const toggleSelect = (id: string) => {
    const date = getDate(id);
    if (!date) return;

    if (selectionDate && date !== selectionDate) {
      // 別日付 → 選択をリセットしてこのアイテムだけ選択
      setSelectedIds(new Set([id]));
      setAnchorId(id);
      setSelectionDate(date);
      return;
    }

    const next = new Set(selectedIds);
    if (next.has(id)) {
      next.delete(id);
    } else {
      next.add(id);
    }

    setSelectedIds(next);
    setAnchorId(id);
    if (!selectionDate) setSelectionDate(date);

    // 選択モード中に最後の1件を解除したら自動的に選択モードを終了する
    if (next.size === 0 && isSelectionMode) {
      setIsSelectionMode(false);
      setAnchorId(null);
      setSelectionDate(null);
    }
  };

  const rangeSelect = (id: string) => {
    const date = getDate(id);
    if (!date) return;

    if (!anchorId || anchorId === id) {
      setAnchorId(id);
      setSelectedIds(new Set([id]));
      setSelectionDate(date);
      return;
    }

    const anchorDate = getDate(anchorId);

    if (!anchorDate || date !== anchorDate) {
      // アンカーと別日付 → 単体選択に切り替え
      setAnchorId(id);
      setSelectedIds(new Set([id]));
      setSelectionDate(date);
      return;
    }

    // 同日付のスケジュール ID だけで範囲を絞る
    const sameDateIds = allSchedules.filter((s) => s.date === anchorDate).map((s) => s.id);
    const anchorIdx = sameDateIds.indexOf(anchorId);
    const targetIdx = sameDateIds.indexOf(id);

    if (anchorIdx === -1 || targetIdx === -1) {
      setAnchorId(id);
      setSelectedIds(new Set([id]));
      return;
    }

    const start = Math.min(anchorIdx, targetIdx);
    const end = Math.max(anchorIdx, targetIdx);
    setSelectedIds(new Set(sameDateIds.slice(start, end + 1)));
    // anchorId は変えない（続けて Shift+Click できるように）
  };

  const clearSelection = () => {
    setSelectedIds(new Set());
    setAnchorId(null);
    setSelectionDate(null);
    setIsSelectionMode(false);
  };

  const isSelected = (id: string) => selectedIds.has(id);

  const enterSelectionMode = (firstId: string) => {
    const date = getDate(firstId);
    if (!date) return;
    setIsSelectionMode(true);
    setSelectedIds(new Set([firstId]));
    setAnchorId(firstId);
    setSelectionDate(date);
  };

  const exitSelectionMode = () => {
    clearSelection();
  };

  return (
    <SelectionContext.Provider value={{ selectedIds, toggleSelect, rangeSelect, clearSelection, isSelected, setAllSchedules, isSelectionMode, enterSelectionMode, exitSelectionMode }}>
      {children}
    </SelectionContext.Provider>
  );
};

export const useSelection = () => {
  const ctx = useContext(SelectionContext);
  if (!ctx) throw new Error('useSelection must be used within SelectionProvider');
  return ctx;
};
