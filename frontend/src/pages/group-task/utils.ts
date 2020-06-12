import * as taskAPI from "../../apis/task";
import { Task, UserTaskStatus } from "../../types";

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

export const normalizeDate = (dateStr?: string): string => {
  if (!dateStr) return new Date().toISOString().slice(0, -5);

  return new Date(dateStr).toISOString().slice(0, -5);
};
