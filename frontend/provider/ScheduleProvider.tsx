'use client';

import { createContext, useContext } from 'react';

import type { Type } from '@/type';
import { Model } from '@/model';
import { useSchedule as useBaseSchedule } from '@/component/calendar/useSchedule';

type ScheduleContextType = {
  masterSchedules: Type.ScheduleDateItem[];
  customSchedules: Type.ScheduleDateItem[];
  setMasterSchedules: (schedules: Type.ScheduleDateItem[]) => void;
  setCustomSchedules: (schedules: Type.ScheduleDateItem[]) => void;
  schedulesByType: (type: Type.ScheduleType) => Type.ScheduleDateItem[];
  setSchedulesByType: (type: Type.ScheduleType, schedules: Type.ScheduleDateItem[]) => void;
  addSchedule: (scheduleRequest: Model.CreateSchedule[]) => Promise<void>;
  removeSchedule: (id: string, type: Type.ScheduleType) => Promise<void>;
  saveSchedule: (scheduleRequest: Model.EditSchedule) => Promise<void>;
  changeScheduleColor: (id: string, type: Type.ScheduleType, color: string) => Promise<void>;
};

const createCtx = () => {
  const ctx = createContext<ScheduleContextType | undefined>(undefined);
  const useCtx = () => {
    const c = useContext(ctx);
    if (!c) throw new Error('useCtx must be inside a Provider with a value');
    return c;
  };
  return [useCtx, ctx.Provider] as const;
};

const [useCtx, SetScheduleProvider] = createCtx();
export const useSchedule = useCtx;

type Props = {
  children: React.ReactNode;
};

export const ScheduleProvider = ({ children }: Props) => {
  const {
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
  } = useBaseSchedule();
  return (
    <SetScheduleProvider
      value={{
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
      }}
    >
      {children}
    </SetScheduleProvider>
  );
};
