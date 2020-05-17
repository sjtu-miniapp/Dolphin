import { Group, Task } from "src/types";

export interface GroupProps {}

export interface GroupState {
  groups: Group[];
}

export type ViewStatus = "Full" | "Short";

export interface TaskViewProps {
  tasks: Task[];
  onClickTask: (taskID: string) => void;
  selectedGroupName?: string;
}

export interface FullGroupViewProps {
  groups: Group[];
  onClickGroup: (groupID: string) => void;
}

export interface GroupTaskViewProps {
  groups: Group[];
  onClickGroup: (groupID: string) => void;
  selectedGroup?: Group;
  tasks: Task[];
  onClickTask: (taskID: string) => void;
}
