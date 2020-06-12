import * as groupAPI from "../../apis/group";
import * as taskAPI from "../../apis/task";
import { Group, Task, UserTaskStatus } from "../../types";

export const normalizeGroup = async (
  group: groupAPI.GroupShort
): Promise<Group> => {
  const { id, name } = group;

  const tasks = await taskAPI.getTasksByGroupID(id);
  const taskNumber = tasks ? tasks.length : 0;
  const updateTime = new Date();

  return { id, name, taskNumber, updateTime };
};

export const normalizeTask = async (
  taskMeta: taskAPI.TaskMeta
): Promise<Task> => {
  const receivers: taskAPI.TaskWorker[] = [];

  try {
    const receiversFromRemoteCall = await taskAPI.getTaskWorker(taskMeta.id);
    receivers.push(...receiversFromRemoteCall);
  } catch (error) {
    console.error(`Failed to get task workers on: ${error}`);
  }

  return {
    id: taskMeta.id,
    groupID: taskMeta.group_id,
    name: taskMeta.name,
    description: taskMeta.description,
    startDate: new Date(taskMeta.start_date),
    endDate: new Date(taskMeta.end_date),
    readOnly: taskMeta.readonly,
    status: taskMeta.done ? "Done" : "Undone",
    receivers: receivers.map(r => normalizeTaskWorker(r))
  };
};

export const normalizeTaskWorker = (
  taskWorker: taskAPI.TaskWorker
): UserTaskStatus => {
  return {
    userName: taskWorker.name,
    status: taskWorker.done ? "completed" : "in-progress"
  };
};
