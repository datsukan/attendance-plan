export class CreateSchedule {
  private name: string;
  private startDate: Date;
  private endDate: Date;
  private color: string;
  private type: 'master' | 'custom';
  private hasBulk: boolean;
  private bulkFrom: number;
  private bulkTo: number;

  constructor(
    name: string,
    startDate: Date,
    endDate: Date,
    color: string,
    type: 'master' | 'custom',
    hasBulk: boolean = false,
    bulkFrom: number = 0,
    bulkTo: number = 0
  ) {
    this.name = name;
    this.startDate = startDate;
    this.endDate = endDate;
    this.color = color;
    this.type = type;
    this.hasBulk = hasBulk;
    this.bulkFrom = bulkFrom;
    this.bulkTo = bulkTo;
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

  public getHasBulk(): boolean {
    return this.hasBulk;
  }

  public getBulkFrom(): number {
    return this.bulkFrom;
  }

  public getBulkTo(): number {
    return this.bulkTo;
  }

  public setBulkFrom(bulkFrom: number): void {
    this.bulkFrom = bulkFrom;
  }

  public setBulkTo(bulkTo: number): void {
    this.bulkTo = bulkTo;
  }
}
