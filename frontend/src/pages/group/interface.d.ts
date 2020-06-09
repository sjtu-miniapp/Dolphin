import { Group, Task } from "src/types";

export interface GroupState {
  groups: Group[];
}

export type ViewStatus = "Full" | "Short";

export interface TaskViewProps {
  tasks: Task[];
  selectedGroupName?: string;
}
export interface GroupTaskViewProps {
  groups: Group[];
  onClickGroup: (groupID: string) => void;
  selectedGroup?: Group;
  tasks: Task[];
  onClickTask: (taskID: string) => void;
}
