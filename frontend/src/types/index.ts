export interface UserTaskStatus {
  userName: string;
  status: string;
}

export interface Task {
  id: number;
  groupID: number;
  name: string;
  description: string;
  publisher?: string;
  leader?: string;
  startDate: Date;
  endDate: Date;
  readOnly: boolean;
  status: string;
  receivers: UserTaskStatus[];
  type?: string;
}
export interface Group {
  id: number;
  name: string;
  taskNumber: number;
  updateTime: Date;
  picUrl?: string;
}
