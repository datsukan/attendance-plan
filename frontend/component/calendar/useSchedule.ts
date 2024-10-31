import { useState } from 'react';

import type { Schedule } from '@/type/schedule';
import { CreateSchedule } from '@/model/create-schedule';
import { EditSchedule } from '@/model/edit-schedule';

export const useSchedule = () => {
  const [schedules, setSchedules] = useState<Schedule[]>([]);

  const addSchedule = (createSchedule: CreateSchedule) => {
    const afterSchedules: Schedule[] = [...schedules];
    if (createSchedule.getHasBulk()) {
      for (let number = createSchedule.getBulkFrom(); number <= createSchedule.getBulkTo(); number++) {
        afterSchedules.push({
          id: createSchedule.getId() + '-' + number,
          name: `第${number}回 ${createSchedule.getName()}`,
          startDate: createSchedule.getStartDate(),
          endDate: createSchedule.getEndDate(),
          color: createSchedule.getColor(),
          type: createSchedule.getType(),
        });
      }
    } else {
      afterSchedules.push({
        id: createSchedule.getId(),
        name: createSchedule.getName(),
        startDate: createSchedule.getStartDate(),
        endDate: createSchedule.getEndDate(),
        color: createSchedule.getColor(),
        type: createSchedule.getType(),
      });
    }

    setSchedules(afterSchedules);
  };

  const removeSchedule = (id: string) => {
    setSchedules(schedules.filter((schedule) => schedule.id !== id));
  };

  const saveSchedule = (editSchedule: EditSchedule) => {
    const set = (schedule: Schedule) => {
      if (schedule.id !== editSchedule.getId()) {
        return schedule;
      }

      schedule.name = editSchedule.getName();
      schedule.startDate = editSchedule.getStartDate();
      schedule.endDate = editSchedule.getEndDate();
      schedule.color = editSchedule.getColor();
      schedule.type = editSchedule.getType();
      return schedule;
    };

    setSchedules(schedules.map(set));
  };

  const changeScheduleColor = (id: string, color: string) => {
    const set = (schedule: Schedule) => {
      if (schedule.id !== id) {
        return schedule;
      }

      schedule.color = color;
      return schedule;
    };

    setSchedules(schedules.map(set));
  };

  return { schedules, setSchedules, addSchedule, removeSchedule, saveSchedule, changeScheduleColor };
};
