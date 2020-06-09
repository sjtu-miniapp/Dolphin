import * as taskAPI from "../../apis/task";
import { Task, UserTaskStatus } from "../../types";

export const normalizeTask = async (
  taskMeta: taskAPI.TaskMeta,
  i: number
): Promise<Task> => {
  const id = i + 1; // FIX ME: no id return from api

  const receivers: taskAPI.TaskWorker[] = [];

  try {
    const receiversFromRemoteCall = await taskAPI.getTaskWorker(id);
    receivers.push(...receiversFromRemoteCall);
  } catch (error) {
    console.error(`Failed to get task workers on: ${error}`);
  }

  return {
    id,
    groupID: taskMeta.groupId,
    name: taskMeta.name,
    description: taskMeta.description,
    startDate: new Date(),
    endDate: new Date(),
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
