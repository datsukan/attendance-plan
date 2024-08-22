import { Schedule } from '@/type/schedule';

export class EditSchedule {
  private id: string;
  private name: string;
  private startDate: Date;
  private endDate: Date;
  private color: string;
  private type: 'master' | 'custom';

  constructor(schedule: Schedule) {
    this.id = schedule.id;
    this.name = schedule.name;
    this.startDate = schedule.startDate;
    this.endDate = schedule.endDate;
    this.color = schedule.color;
    this.type = schedule.type;
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
    return this.type;
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
    this.type = type;
  }
}
