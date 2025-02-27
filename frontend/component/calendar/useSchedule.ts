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

export const useSchedule = () => {
  const [masterSchedules, setMasterSchedules] = useState<Type.ScheduleDateItem[]>([]);
  const [customSchedules, setCustomSchedules] = useState<Type.ScheduleDateItem[]>([]);

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

  const addSchedule = async (reqSchedules: Model.CreateSchedule[]): Promise<void> => {
    if (reqSchedules.length === 0) {
      return;
    }

    const targetSchedules = schedulesByType(reqSchedules[0].getType());
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

    try {
      const resSchedules = await createBulkSchedules(createBulkScheduleParams);
      const dateKey = Model.ScheduleDateItem.toKey(createBulkScheduleParams[0].startDate);
      const type = new Model.ScheduleType(createBulkScheduleParams[0].type);
      for (const resSchedule of resSchedules) {
        resultSchedules.addSchedule(dateKey, type, new Schedule(resSchedule));
      }
    } catch (e) {
      toast.error(String(e));
      return;
    }

    setSchedulesByType(reqSchedules[0].getType(), resultSchedules.toTypeScheduleDateItems());
  };

  const removeSchedule = async (id: string, type: Type.ScheduleType): Promise<void> => {
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

    try {
      await updateSchedule(s);
    } catch (e) {
      toast.error(String(e));
      return;
    }

    const allSchedules = [...masterSchedules, ...customSchedules];
    const beforeSchedule = new Model.ScheduleDateItemList(allSchedules).getSchedule(s.id);
    if (!beforeSchedule) {
      return;
    }

    if (beforeSchedule.getType() === s.type) {
      const typeSchedules = schedulesByType(s.type);
      const typeResultSchedules = new Model.ScheduleDateItemList(typeSchedules);
      typeResultSchedules.updateSchedule(new Model.Schedule(s));
      setSchedulesByType(s.type, typeResultSchedules.toTypeScheduleDateItems());
      return;
    }

    const beforeTypeSchedules = schedulesByType(beforeSchedule.getType());
    const beforeTypeResultSchedules = new Model.ScheduleDateItemList(beforeTypeSchedules);
    beforeTypeResultSchedules.removeSchedule(s.id);

    setSchedulesByType(beforeSchedule.getType(), beforeTypeResultSchedules.toTypeScheduleDateItems());

    const afterTypeSchedules = schedulesByType(s.type);
    const afterTypeResultSchedules = new Model.ScheduleDateItemList(afterTypeSchedules);
    afterTypeResultSchedules.addSchedule(Model.ScheduleDateItem.toKey(s.startDate), new Model.ScheduleType(s.type), new Model.Schedule(s));

    setSchedulesByType(s.type, afterTypeResultSchedules.toTypeScheduleDateItems());
  };

  const changeScheduleColor = async (id: string, type: Type.ScheduleType, color: string): Promise<void> => {
    const targetSchedules = schedulesByType(type);
    const resultSchedules = new Model.ScheduleDateItemList(targetSchedules);
    const schedule = resultSchedules.getSchedule(id);
    if (!schedule) {
      return;
    }

    schedule.setColor(color);

    try {
      await updateSchedule(schedule.toTypeSchedule());
    } catch (e) {
      toast.error(String(e));
      return;
    }

    resultSchedules.updateSchedule(schedule);
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
    saveSchedule,
    changeScheduleColor,
  };
};
