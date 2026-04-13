import { useEffect, useState } from 'react';
import { toast } from 'react-hot-toast';

import type { Type } from '@/type';
import { ScheduleTypeMaster } from '@/const/schedule';
import { Model } from '@/model';

import { fetchSchedules } from '@/backend-api/fetchSchedules';
import { createBulkSchedules, CreateScheduleParam } from '@/backend-api/createBulkSchedules';
import { updateSchedule } from '@/backend-api/updateSchedule';
import { deleteSchedule } from '@/backend-api/deleteSchedule';
import { Schedule } from '@/model/schedule';
import { useUndo } from '@/provider/UndoProvider';

export const useSchedule = () => {
  const [masterSchedules, setMasterSchedules] = useState<Type.ScheduleDateItem[]>([]);
  const [customSchedules, setCustomSchedules] = useState<Type.ScheduleDateItem[]>([]);
  const { setUndoCommand } = useUndo();

  useEffect(() => {
    (async () => {
      try {
        const result = await fetchSchedules();
        setMasterSchedules(result.masterSchedules);
        setCustomSchedules(result.customSchedules);
      } catch (e) {
        toast.error(String(e));
      }
    })();
  }, []);

  const schedulesByType = (type: Type.ScheduleType): Type.ScheduleDateItem[] => {
    return type === ScheduleTypeMaster ? masterSchedules : customSchedules;
  };

  const setSchedulesByType = (type: Type.ScheduleType, schedules: Type.ScheduleDateItem[]) => {
    if (type === ScheduleTypeMaster) {
      setMasterSchedules(schedules);
      return;
    }

    setCustomSchedules(schedules);
  };

  const setSchedulesByTypeFunctional = (
    type: Type.ScheduleType,
    updater: (prev: Type.ScheduleDateItem[]) => Type.ScheduleDateItem[]
  ) => {
    if (type === ScheduleTypeMaster) {
      setMasterSchedules(updater);
      return;
    }

    setCustomSchedules(updater);
  };

  const addSchedule = async (reqSchedules: Model.CreateSchedule[]): Promise<void> => {
    if (reqSchedules.length === 0) {
      return;
    }

    const scheduleType = reqSchedules[0].getType();
    const targetSchedules = schedulesByType(scheduleType);
    const resultSchedules = new Model.ScheduleDateItemList(targetSchedules);

    let createBulkScheduleParams: CreateScheduleParam[] = [];
    for (const reqSchedule of reqSchedules) {
      const s: CreateScheduleParam = {
        name: reqSchedule.getName(),
        startDate: reqSchedule.getStartDate(),
        endDate: reqSchedule.getEndDate(),
        color: reqSchedule.getColor(),
        type: reqSchedule.getType(),
      };

      createBulkScheduleParams.push(s);
    }

    let resSchedules: Type.Schedule[];
    try {
      resSchedules = await createBulkSchedules(createBulkScheduleParams);
      const dateKey = Model.ScheduleDateItem.toKey(createBulkScheduleParams[0].startDate);
      const type = new Model.ScheduleType(createBulkScheduleParams[0].type);
      for (const resSchedule of resSchedules) {
        resultSchedules.addSchedule(dateKey, type, new Schedule(resSchedule));
      }
    } catch (e) {
      toast.error(String(e));
      return;
    }

    setSchedulesByType(scheduleType, resultSchedules.toTypeScheduleDateItems());

    const createdIds = resSchedules.map((s) => s.id);
    setUndoCommand({
      label: `スケジュールを${resSchedules.length}件追加しました`,
      execute: async () => {
        await Promise.all(createdIds.map((id) => deleteSchedule(id)));
        setSchedulesByTypeFunctional(scheduleType, (prev) => {
          const list = new Model.ScheduleDateItemList(prev);
          createdIds.forEach((id) => list.removeSchedule(id));
          return list.toTypeScheduleDateItems();
        });
      },
    });
  };

  const removeSchedule = async (id: string, type: Type.ScheduleType): Promise<void> => {
    // Capture schedule data before deletion for undo
    const capturedList = new Model.ScheduleDateItemList(schedulesByType(type));
    const capturedModel = capturedList.getSchedule(id);
    const capturedData: Type.Schedule | undefined = capturedModel?.toTypeSchedule();

    try {
      await deleteSchedule(id);
    } catch (e) {
      toast.error(String(e));
      return;
    }

    const targetSchedules = schedulesByType(type);
    const resultSchedules = new Model.ScheduleDateItemList(targetSchedules);
    resultSchedules.removeSchedule(id);
    setSchedulesByType(type, resultSchedules.toTypeScheduleDateItems());

    if (!capturedData) {
      return;
    }

    setUndoCommand({
      label: `「${capturedData.name}」を削除しました`,
      execute: async () => {
        const recreated = await createBulkSchedules([
          {
            name: capturedData.name,
            startDate: capturedData.startDate,
            endDate: capturedData.endDate,
            color: capturedData.color,
            type: capturedData.type,
          },
        ]);
        if (recreated.length === 0) return;
        setSchedulesByTypeFunctional(type, (prev) => {
          const list = new Model.ScheduleDateItemList(prev);
          const dateKey = Model.ScheduleDateItem.toKey(recreated[0].startDate);
          const schedType = new Model.ScheduleType(recreated[0].type);
          list.addSchedule(dateKey, schedType, new Schedule(recreated[0]));
          return list.toTypeScheduleDateItems();
        });
      },
    });
  };

  const removeBulkSchedules = async (schedules: Type.Schedule[]): Promise<void> => {
    // Snapshot before deletion for undo
    const capturedSchedules = schedules.map((s) => ({ ...s }));

    try {
      await Promise.all(schedules.map((s) => deleteSchedule(s.id)));
    } catch (e) {
      toast.error(String(e));
      return;
    }

    const masterIds = schedules.filter((s) => s.type === ScheduleTypeMaster).map((s) => s.id);
    const customIds = schedules.filter((s) => s.type !== ScheduleTypeMaster).map((s) => s.id);

    if (masterIds.length > 0) {
      const result = new Model.ScheduleDateItemList(masterSchedules);
      masterIds.forEach((id) => result.removeSchedule(id));
      setMasterSchedules(result.toTypeScheduleDateItems());
    }

    if (customIds.length > 0) {
      const result = new Model.ScheduleDateItemList(customSchedules);
      customIds.forEach((id) => result.removeSchedule(id));
      setCustomSchedules(result.toTypeScheduleDateItems());
    }

    const toParam = (s: Type.Schedule): CreateScheduleParam => ({
      name: s.name,
      startDate: s.startDate,
      endDate: s.endDate,
      color: s.color,
      type: s.type,
    });

    setUndoCommand({
      label: `スケジュールを${capturedSchedules.length}件削除しました`,
      execute: async () => {
        const masterToRecreate = capturedSchedules.filter((s) => s.type === ScheduleTypeMaster);
        const customToRecreate = capturedSchedules.filter((s) => s.type !== ScheduleTypeMaster);

        if (masterToRecreate.length > 0) {
          const recreated = await createBulkSchedules(masterToRecreate.map(toParam));
          setMasterSchedules((prev) => {
            const list = new Model.ScheduleDateItemList(prev);
            for (const s of recreated) {
              list.addSchedule(
                Model.ScheduleDateItem.toKey(s.startDate),
                new Model.ScheduleType(s.type),
                new Schedule(s)
              );
            }
            return list.toTypeScheduleDateItems();
          });
        }

        if (customToRecreate.length > 0) {
          const recreated = await createBulkSchedules(customToRecreate.map(toParam));
          setCustomSchedules((prev) => {
            const list = new Model.ScheduleDateItemList(prev);
            for (const s of recreated) {
              list.addSchedule(
                Model.ScheduleDateItem.toKey(s.startDate),
                new Model.ScheduleType(s.type),
                new Schedule(s)
              );
            }
            return list.toTypeScheduleDateItems();
          });
        }
      },
    });
  };

  const saveSchedule = async (scheduleRequest: Model.EditSchedule): Promise<void> => {
    const s: Type.Schedule = {
      id: scheduleRequest.getId(),
      name: scheduleRequest.getName(),
      startDate: scheduleRequest.getStartDate(),
      endDate: scheduleRequest.getEndDate(),
      color: scheduleRequest.getColor(),
      type: scheduleRequest.getType(),
      order: scheduleRequest.getOrder(),
    };

    // Capture before state for undo (must be before API call)
    const allSchedules = [...masterSchedules, ...customSchedules];
    const beforeModel = new Model.ScheduleDateItemList(allSchedules).getSchedule(s.id);
    const beforeData: Type.Schedule | undefined = beforeModel?.toTypeSchedule();

    if (!beforeData) {
      return;
    }

    try {
      await updateSchedule(s);
    } catch (e) {
      toast.error(String(e));
      return;
    }

    if (beforeData.type === s.type) {
      const typeSchedules = schedulesByType(s.type);
      const typeResultSchedules = new Model.ScheduleDateItemList(typeSchedules);
      typeResultSchedules.updateSchedule(new Model.Schedule(s));
      setSchedulesByType(s.type, typeResultSchedules.toTypeScheduleDateItems());
    } else {
      const beforeTypeSchedules = schedulesByType(beforeData.type);
      const beforeTypeResultSchedules = new Model.ScheduleDateItemList(beforeTypeSchedules);
      beforeTypeResultSchedules.removeSchedule(s.id);
      setSchedulesByType(beforeData.type, beforeTypeResultSchedules.toTypeScheduleDateItems());

      const afterTypeSchedules = schedulesByType(s.type);
      const afterTypeResultSchedules = new Model.ScheduleDateItemList(afterTypeSchedules);
      afterTypeResultSchedules.addSchedule(
        Model.ScheduleDateItem.toKey(s.startDate),
        new Model.ScheduleType(s.type),
        new Model.Schedule(s)
      );
      setSchedulesByType(s.type, afterTypeResultSchedules.toTypeScheduleDateItems());
    }

    setUndoCommand({
      label: `「${s.name}」を編集しました`,
      execute: async () => {
        await updateSchedule(beforeData);

        if (beforeData.type === s.type) {
          setSchedulesByTypeFunctional(beforeData.type, (prev) => {
            const list = new Model.ScheduleDateItemList(prev);
            list.updateSchedule(new Model.Schedule(beforeData));
            return list.toTypeScheduleDateItems();
          });
        } else {
          // Revert type change: remove from after-type, restore to before-type
          setSchedulesByTypeFunctional(s.type, (prev) => {
            const list = new Model.ScheduleDateItemList(prev);
            list.removeSchedule(beforeData.id);
            return list.toTypeScheduleDateItems();
          });
          setSchedulesByTypeFunctional(beforeData.type, (prev) => {
            const list = new Model.ScheduleDateItemList(prev);
            list.addSchedule(
              Model.ScheduleDateItem.toKey(beforeData.startDate),
              new Model.ScheduleType(beforeData.type),
              new Model.Schedule(beforeData)
            );
            return list.toTypeScheduleDateItems();
          });
        }
      },
    });
  };

  const changeScheduleColor = async (id: string, type: Type.ScheduleType, color: string): Promise<void> => {
    const targetSchedules = schedulesByType(type);
    const resultSchedules = new Model.ScheduleDateItemList(targetSchedules);
    const schedule = resultSchedules.getSchedule(id);
    if (!schedule) {
      return;
    }

    // Capture before color change for undo
    const scheduleName = schedule.getName();
    const scheduleBeforeChange = schedule.toTypeSchedule();

    schedule.setColor(color);

    try {
      await updateSchedule(schedule.toTypeSchedule());
    } catch (e) {
      toast.error(String(e));
      return;
    }

    resultSchedules.updateSchedule(schedule);
    setSchedulesByType(type, resultSchedules.toTypeScheduleDateItems());

    setUndoCommand({
      label: `「${scheduleName}」のカラーを変更しました`,
      execute: async () => {
        await updateSchedule(scheduleBeforeChange);
        setSchedulesByTypeFunctional(type, (prev) => {
          const list = new Model.ScheduleDateItemList(prev);
          const s = list.getSchedule(id);
          if (!s) return prev;
          s.setColor(scheduleBeforeChange.color);
          list.updateSchedule(s);
          return list.toTypeScheduleDateItems();
        });
      },
    });
  };

  return {
    masterSchedules,
    customSchedules,
    setMasterSchedules,
    setCustomSchedules,
    schedulesByType,
    setSchedulesByType,
    addSchedule,
    removeSchedule,
    removeBulkSchedules,
    saveSchedule,
    changeScheduleColor,
  };
};
