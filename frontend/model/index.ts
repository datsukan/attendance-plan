import { Schedule as M_Schedule, ScheduleType as M_ScheduleType } from './schedule';
import { ScheduleDateItem as M_ScheduleDateItem, ScheduleDateItemList as M_ScheduleDateItemList } from './scheduleDateItem';
import { CreateSchedule as M_CreateSchedule } from './createSchedule';
import { EditSchedule as M_EditSchedule } from './editSchedule';

export namespace Model {
  export class Schedule extends M_Schedule {}
  export class ScheduleType extends M_ScheduleType {}
  export class ScheduleDateItem extends M_ScheduleDateItem {}
  export class ScheduleDateItemList extends M_ScheduleDateItemList {}
  export class CreateSchedule extends M_CreateSchedule {}
  export class EditSchedule extends M_EditSchedule {}
}
