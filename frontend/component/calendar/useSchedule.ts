import { useEffect, useState } from 'react';
import { toast } from 'react-hot-toast';

import type { Type } from '@/type';
import { ScheduleTypeMaster } from '@/const/schedule';
import { Model } from '@/model';

import { fetchSchedules } from '@/backend-api/fetchSchedules';
import { createSchedule, CreateScheduleParam } from '@/backend-api/createSchedule';
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

  const addSchedule = async (scheduleRequest: Model.CreateSchedule): Promise<void> => {
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

        try {
          const resSchedule = await createSchedule(s);
          const dateKey = Model.ScheduleDateItem.toKey(s.startDate);
          const type = new Model.ScheduleType(s.type);
          resultSchedules.addSchedule(dateKey, type, new Schedule(resSchedule));
        } catch (e) {
          toast.error(String(e));
          return;
        }
      }
    } else {
      const s = {
        name: scheduleRequest.getName(),
        startDate: scheduleRequest.getStartDate(),
        endDate: scheduleRequest.getEndDate(),
        color: scheduleRequest.getColor(),
        type: scheduleRequest.getType(),
      };

      try {
        const resSchedule = await createSchedule(s);
        const dateKey = Model.ScheduleDateItem.toKey(s.startDate);
        const type = new Model.ScheduleType(s.type);
        resultSchedules.addSchedule(dateKey, type, new Schedule(resSchedule));
      } catch (e) {
        toast.error(String(e));
        return;
      }
    }

    setSchedulesByType(scheduleRequest.getType(), resultSchedules.toTypeScheduleDateItems());
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
