import { useMemo, useCallback } from 'react';
import { arrayMove } from '@dnd-kit/sortable';

import { Type } from '@/type';
import { Model } from '@/model';
import { useSchedule } from '@/provider/ScheduleProvider';
import { shiftSchedule } from './drag-helpers';

export type ScheduleMove = {
  schedule: Type.Schedule;
  newStartDate: Date;
  newType: Type.ScheduleType;
};

/**
 * スケジュール状態への読み書きを抽象化するフック。
 *
 * 呼び出し側は ScheduleDateItemList の構築を意識しなくてよい。
 * 内部では useMemo により master/custom の結合リストをキャッシュしている。
 */
export const useScheduleStore = () => {
  const {
    masterSchedules,
    customSchedules,
    setSchedulesByType,
    setSchedulesByTypeFunctional,
  } = useSchedule();

  // 毎レンダーの再構築を避けるためメモ化する
  const allList = useMemo(
    () => new Model.ScheduleDateItemList([...masterSchedules, ...customSchedules]),
    [masterSchedules, customSchedules]
  );

  // ─── 参照 ──────────────────────────────────────────────────────────────────

  const findById = useCallback(
    (id: string): Type.Schedule | undefined =>
      allList.getSchedule(id)?.toTypeSchedule(),
    [allList]
  );

  const getCell = useCallback(
    (dateKey: string, type: Type.ScheduleType): Type.Schedule[] => {
      const mType = new Model.ScheduleType(type);
      return (
        allList.getDateItem(dateKey, mType)?.toTypeSchedules() ?? []
      );
    },
    [allList]
  );

  // ─── 変更：1件移動（dragOver 用の楽観的更新） ─────────────────────────────

  const applyMove = useCallback(
    (move: ScheduleMove): void => {
      const { schedule, newStartDate, newType } = move;
      const movedSchedule = { ...shiftSchedule(schedule, newStartDate), type: newType };

      if (schedule.type === newType) {
        // 同タイプ内移動
        const list = new Model.ScheduleDateItemList(
          newType === 'master' ? [...masterSchedules] : [...customSchedules]
        );
        list.removeSchedule(schedule.id);
        list.addSchedule(
          Model.ScheduleDateItem.toKey(newStartDate),
          new Model.ScheduleType(newType),
          new Model.Schedule(movedSchedule)
        );
        setSchedulesByType(newType, list.toTypeScheduleDateItems());
      } else {
        // タイプ変更を伴う移動
        const beforeList = new Model.ScheduleDateItemList(
          schedule.type === 'master' ? [...masterSchedules] : [...customSchedules]
        );
        const afterList = new Model.ScheduleDateItemList(
          newType === 'master' ? [...masterSchedules] : [...customSchedules]
        );
        beforeList.removeSchedule(schedule.id);
        afterList.addSchedule(
          Model.ScheduleDateItem.toKey(newStartDate),
          new Model.ScheduleType(newType),
          new Model.Schedule({ ...movedSchedule, order: 0 })
        );
        setSchedulesByType(schedule.type, beforeList.toTypeScheduleDateItems());
        setSchedulesByType(newType, afterList.toTypeScheduleDateItems());
      }
    },
    [masterSchedules, customSchedules, setSchedulesByType]
  );

  // ─── 変更：複数件移動（dragEnd 用） ────────────────────────────────────────

  const applyMoves = useCallback(
    (moves: ScheduleMove[]): void => {
      // 移動後スケジュールを構築
      const movedItems: Type.Schedule[] = moves.map(({ schedule, newStartDate, newType }) => ({
        ...shiftSchedule(schedule, newStartDate),
        type: newType,
      }));

      // master・custom 両リストから対象スケジュールを削除
      const masterList = new Model.ScheduleDateItemList([...masterSchedules]);
      const customList = new Model.ScheduleDateItemList([...customSchedules]);
      for (const { schedule } of moves) {
        masterList.removeSchedule(schedule.id);
        customList.removeSchedule(schedule.id);
      }

      // 各移動先のリストへ追加
      for (const moved of movedItems) {
        const targetList = moved.type === 'master' ? masterList : customList;
        targetList.addSchedule(
          Model.ScheduleDateItem.toKey(moved.startDate),
          new Model.ScheduleType(moved.type),
          new Model.Schedule(moved)
        );
      }

      setSchedulesByType('master', masterList.toTypeScheduleDateItems());
      setSchedulesByType('custom', customList.toTypeScheduleDateItems());
    },
    [masterSchedules, customSchedules, setSchedulesByType]
  );

  // ─── 変更：同セル内並び替え ─────────────────────────────────────────────────

  const reorderCell = useCallback(
    (
      dateKey: string,
      type: Type.ScheduleType,
      fromIndex: number,
      toIndex: number
    ): Type.Schedule[] | null => {
      const list = new Model.ScheduleDateItemList(
        type === 'master' ? [...masterSchedules] : [...customSchedules]
      );
      const mType = new Model.ScheduleType(type);
      const dateItem = list.getDateItem(dateKey, mType);
      if (!dateItem) return null;

      const schedules = dateItem.getSchedules();
      if (
        fromIndex < 0 ||
        toIndex < 0 ||
        fromIndex >= schedules.length ||
        toIndex >= schedules.length
      ) return null;

      const reordered = arrayMove(schedules, fromIndex, toIndex);
      reordered.forEach((s, i) => s.setOrder(i + 1));
      list.setSchedules(dateKey, mType, reordered);
      setSchedulesByType(type, list.toTypeScheduleDateItems());
      // setState は非同期のため、並び替え後データを呼び出し元へ直接返す
      return reordered.map((s) => s.toTypeSchedule());
    },
    [masterSchedules, customSchedules, setSchedulesByType]
  );

  // ─── 変更：スナップショットからの復元（Undo 用） ───────────────────────────

  const restoreSnapshot = useCallback(
    (snapshot: Type.Schedule[]): void => {
      const snapshotIds = new Set(snapshot.map((s) => s.id));
      const masterOriginals = snapshot.filter((s) => s.type === 'master');
      const customOriginals = snapshot.filter((s) => s.type !== 'master');

      setSchedulesByTypeFunctional('master', (prev) => {
        const list = new Model.ScheduleDateItemList(prev);
        snapshotIds.forEach((id) => list.removeSchedule(id));
        masterOriginals.forEach((s) =>
          list.addSchedule(
            Model.ScheduleDateItem.toKey(s.startDate),
            new Model.ScheduleType(s.type),
            new Model.Schedule(s)
          )
        );
        return list.toTypeScheduleDateItems();
      });

      setSchedulesByTypeFunctional('custom', (prev) => {
        const list = new Model.ScheduleDateItemList(prev);
        snapshotIds.forEach((id) => list.removeSchedule(id));
        customOriginals.forEach((s) =>
          list.addSchedule(
            Model.ScheduleDateItem.toKey(s.startDate),
            new Model.ScheduleType(s.type),
            new Model.Schedule(s)
          )
        );
        return list.toTypeScheduleDateItems();
      });
    },
    [setSchedulesByTypeFunctional]
  );

  return {
    findById,
    getCell,
    applyMove,
    applyMoves,
    reorderCell,
    restoreSnapshot,
  };
};
