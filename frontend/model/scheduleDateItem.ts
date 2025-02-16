import { format } from 'date-fns';

import { Type } from '@/type';
import { Schedule, ScheduleType } from './schedule';

export class ScheduleDateItem {
  private static dateFormat: string = 'yyyy-MM-dd';
  private date: string;
  private type: ScheduleType;
  private schedules: Schedule[];

  constructor(scheduleDateItem: Type.ScheduleDateItem, keyDate: Date | null = null) {
    this.date = keyDate ? ScheduleDateItem.toKey(keyDate) : scheduleDateItem.date;
    this.type = new ScheduleType(scheduleDateItem.type);
    this.schedules = scheduleDateItem.schedules.map((schedule) => new Schedule(schedule));
  }

  public static toKey(date: Date): string {
    return format(date, this.dateFormat);
  }

  public getDate(): string {
    return this.date;
  }

  public getType(): 'master' | 'custom' {
    return this.type.String();
  }

  public getSchedules(): Schedule[] {
    return this.schedules;
  }

  public toTypeSchedules(): Type.Schedule[] {
    return this.schedules.map((schedule) => ({
      id: schedule.getId(),
      name: schedule.getName(),
      startDate: schedule.getStartDate(),
      endDate: schedule.getEndDate(),
      color: schedule.getColor(),
      type: schedule.getType(),
      order: schedule.getOrder(),
    }));
  }

  public getLastOrder(): number {
    let max = 0;
    for (const schedule of this.schedules) {
      if (schedule.getOrder() > max) {
        max = schedule.getOrder();
      }
    }

    return max;
  }

  public getNextOrder(): number {
    return this.getLastOrder() + 1;
  }

  public addSchedule(schedule: Schedule): void {
    schedule.setOrder(this.getNextOrder());
    this.schedules.push(schedule);
  }

  public setSchedules(schedules: Schedule[]): void {
    this.schedules = schedules;
  }
}

export class ScheduleDateItemList {
  private scheduleDateItems: ScheduleDateItem[];

  constructor(scheduleDateItems: Type.ScheduleDateItem[]) {
    this.scheduleDateItems = scheduleDateItems.map((scheduleDateItem) => new ScheduleDateItem(scheduleDateItem));
  }

  public getScheduleDateItems(): ScheduleDateItem[] {
    return this.scheduleDateItems;
  }

  public toTypeScheduleDateItems(): Type.ScheduleDateItem[] {
    let typeScheduleDateItems: Type.ScheduleDateItem[] = [];
    for (const scheduleDateItem of this.scheduleDateItems) {
      typeScheduleDateItems.push({
        date: scheduleDateItem.getDate(),
        type: scheduleDateItem.getType(),
        schedules: scheduleDateItem.getSchedules().map((schedule) => ({
          id: schedule.getId(),
          name: schedule.getName(),
          startDate: schedule.getStartDate(),
          endDate: schedule.getEndDate(),
          color: schedule.getColor(),
          type: schedule.getType(),
          order: schedule.getOrder(),
        })),
      });
    }

    return typeScheduleDateItems;
  }

  public getSchedule(id: string): Schedule | undefined {
    for (const scheduleDateItem of this.scheduleDateItems) {
      const schedule = scheduleDateItem.getSchedules().find((schedule) => schedule.getId() === id);
      if (schedule) {
        return schedule;
      }
    }

    return undefined;
  }

  public getDateItem(date: string, type: ScheduleType): ScheduleDateItem | undefined {
    return this.scheduleDateItems.find(
      (scheduleDateItem) => scheduleDateItem.getDate() === date && scheduleDateItem.getType() === type.String()
    );
  }

  public addScheduleDateItem(scheduleDateItem: ScheduleDateItem): void {
    this.scheduleDateItems.push(scheduleDateItem);
  }

  public setSchedules(date: string, type: ScheduleType, schedules: Schedule[]): void {
    const scheduleDateItem = this.getDateItem(date, type);
    if (!scheduleDateItem) {
      return;
    }

    scheduleDateItem.setSchedules(schedules);
  }

  public addSchedule(date: string, type: ScheduleType, schedule: Schedule): void {
    const scheduleDateItem = this.getDateItem(date, type);
    if (!scheduleDateItem) {
      const di = new ScheduleDateItem({ date, type: type.String(), schedules: [schedule.toTypeSchedule()] });
      this.addScheduleDateItem(di);
      return;
    }

    scheduleDateItem.addSchedule(schedule);
  }

  public updateSchedule(schedule: Schedule): void {
    this.removeSchedule(schedule.getId());

    const dateKey = ScheduleDateItem.toKey(schedule.getStartDate());
    const type = new ScheduleType(schedule.getType());
    this.addSchedule(dateKey, type, schedule);
  }

  public removeSchedule(id: string): void {
    for (const scheduleDateItem of this.scheduleDateItems) {
      const schedules = scheduleDateItem.getSchedules().filter((schedule) => schedule.getId() !== id);
      scheduleDateItem.setSchedules(schedules);
    }
  }
}
