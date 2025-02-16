import { useEffect, useState } from 'react';

import type { Type } from '@/type';
import { ScheduleTypeMaster } from '@/const/schedule';
import { Model } from '@/model';

import { fetchSchedules } from '@/api/fetchSchedules';
import { createSchedule, CreateScheduleParam } from '@/api/createSchedule';
import { updateSchedule } from '@/api/updateSchedule';
import { deleteSchedule } from '@/api/deleteSchedule';
import { Schedule } from '@/model/schedule';

export const useSchedule = () => {
  const [masterSchedules, setMasterSchedules] = useState<Type.ScheduleDateItem[]>([]);
  const [customSchedules, setCustomSchedules] = useState<Type.ScheduleDateItem[]>([]);

  useEffect(() => {
    (async () => {
      const result = await fetchSchedules();
      setMasterSchedules(result.masterSchedules);
      setCustomSchedules(result.customSchedules);
    })();
  }, []);

  const schedulesByType = (type: Type.ScheduleType) => {
    return type === ScheduleTypeMaster ? masterSchedules : customSchedules;
  };

  const setSchedulesByType = (type: Type.ScheduleType, schedules: Type.ScheduleDateItem[]) => {
    if (type === ScheduleTypeMaster) {
      setMasterSchedules(schedules);
      return;
    }

    setCustomSchedules(schedules);
  };

  const addSchedule = async (scheduleRequest: Model.CreateSchedule) => {
    const targetSchedules = schedulesByType(scheduleRequest.getType());
    const resultSchedules = new Model.ScheduleDateItemList(targetSchedules);
    if (scheduleRequest.getHasBulk()) {
      for (let number = scheduleRequest.getBulkFrom(); number <= scheduleRequest.getBulkTo(); number++) {
        const s: CreateScheduleParam = {
          name: `第${number}回 ${scheduleRequest.getName()}`,
          startDate: scheduleRequest.getStartDate(),
          endDate: scheduleRequest.getEndDate(),
          color: scheduleRequest.getColor(),
          type: scheduleRequest.getType(),
        };
        const resSchedule = await createSchedule(s);
        if (resSchedule === null) {
          continue;
        }

        const dateKey = Model.ScheduleDateItem.toKey(s.startDate);
        const type = new Model.ScheduleType(s.type);
        resultSchedules.addSchedule(dateKey, type, new Schedule(resSchedule));
      }
    } else {
      const s = {
        name: scheduleRequest.getName(),
        startDate: scheduleRequest.getStartDate(),
        endDate: scheduleRequest.getEndDate(),
        color: scheduleRequest.getColor(),
        type: scheduleRequest.getType(),
      };
      const resSchedule = await createSchedule(s);
      if (resSchedule === null) {
        return;
      }

      const dateKey = Model.ScheduleDateItem.toKey(s.startDate);
      const type = new Model.ScheduleType(s.type);
      resultSchedules.addSchedule(dateKey, type, new Schedule(resSchedule));
    }

    setSchedulesByType(scheduleRequest.getType(), resultSchedules.toTypeScheduleDateItems());
  };

  const removeSchedule = async (id: string, type: Type.ScheduleType) => {
    await deleteSchedule(id);
    const targetSchedules = schedulesByType(type);
    const resultSchedules = new Model.ScheduleDateItemList(targetSchedules);
    resultSchedules.removeSchedule(id);
    setSchedulesByType(type, resultSchedules.toTypeScheduleDateItems());
  };

  const saveSchedule = async (scheduleRequest: Model.EditSchedule) => {
    const s: Type.Schedule = {
      id: scheduleRequest.getId(),
      name: scheduleRequest.getName(),
      startDate: scheduleRequest.getStartDate(),
      endDate: scheduleRequest.getEndDate(),
      color: scheduleRequest.getColor(),
      type: scheduleRequest.getType(),
      order: scheduleRequest.getOrder(),
    };

    await updateSchedule(s);

    const targetSchedules = schedulesByType(s.type);
    const resultSchedules = new Model.ScheduleDateItemList(targetSchedules);
    resultSchedules.updateSchedule(new Schedule(s));

    setSchedulesByType(s.type, resultSchedules.toTypeScheduleDateItems());
  };

  const changeScheduleColor = async (id: string, type: Type.ScheduleType, color: string) => {
    const targetSchedules = schedulesByType(type);
    const resultSchedules = new Model.ScheduleDateItemList(targetSchedules);
    const schedule = resultSchedules.getSchedule(id);
    if (!schedule) {
      return;
    }

    schedule.setColor(color);

    await updateSchedule(schedule.toTypeSchedule());

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
