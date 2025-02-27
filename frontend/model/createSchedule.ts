export class CreateSchedule {
  private name: string;
  private startDate: Date;
  private endDate: Date;
  private color: string;
  private type: 'master' | 'custom';

  constructor(name: string, startDate: Date, endDate: Date, color: string, type: 'master' | 'custom') {
    this.name = name;
    this.startDate = startDate;
    this.endDate = endDate;
    this.color = color;
    this.type = type;
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

  public isMasterType(): boolean {
    return this.type === 'master';
  }

  public isCustomType(): boolean {
    return this.type === 'custom';
  }
}
