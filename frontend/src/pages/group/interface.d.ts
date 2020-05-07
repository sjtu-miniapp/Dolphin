import { Group, Task } from "src/types";

export interface GroupProps {}

export interface GroupState {
  groups: Group[];
}

export type ViewStatus = "Full" | "Short";

export interface GroupViewProps {
  groups: Group[];
  viewStatus: ViewStatus;
  onClickGroup: (groupID: string) => void;
  seletectGroup?: Group;
}

export interface TaskViewProps {
  tasks: Task[];
  onClickTask: (taskID: string) => void;
  selectedGroupName?: string;
}
