import { Type } from '@/type/index';

export class ScheduleType {
  private type: Type.ScheduleType;

  constructor(type: string) {
    if (type !== 'master' && type !== 'custom') {
      this.type = 'master';
      return;
    }
    this.type = type;
  }

  public String(): Type.ScheduleType {
    return this.type;
  }
}

export class Schedule {
  private id: string;
  private name: string;
  private startDate: Date;
  private endDate: Date;
  private color: string;
  private type: ScheduleType;
  private order: number;

  constructor(schedule: Type.Schedule) {
    this.id = schedule.id;
    this.name = schedule.name;
    this.startDate = schedule.startDate;
    this.endDate = schedule.endDate;
    this.color = schedule.color;
    this.type = new ScheduleType(schedule.type);
    this.order = schedule.order;
  }

  public getId(): string {
    return this.id;
  }

  public getName(): string {
    return this.name;
  }

  public getStartDate(): Date {
    return this.startDate;
  }

  public getEndDate(): Date {
    return this.endDate;
  }

  public getColor(): string {
    return this.color;
  }

  public getType(): 'master' | 'custom' {
    return this.type.String();
  }

  public getOrder(): number {
    return this.order;
  }

  public toTypeSchedule(): Type.Schedule {
    return {
      id: this.id,
      name: this.name,
      startDate: this.startDate,
      endDate: this.endDate,
      color: this.color,
      type: this.type.String(),
      order: this.order,
    };
  }

  public setName(name: string): void {
    this.name = name;
  }

  public setStartDate(startDate: Date): void {
    this.startDate = startDate;
  }

  public setEndDate(endDate: Date): void {
    this.endDate = endDate;
  }

  public setColor(color: string): void {
    this.color = color;
  }

  public setType(type: 'master' | 'custom'): void {
    this.type = new ScheduleType(type);
  }

  public setOrder(order: number): void {
    this.order = order;
  }
}
