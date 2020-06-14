import * as taskAPI from "./apis/task";
import { Task, UserTaskStatus } from "./types";

export function cloneDeep(obj: any, hash = new WeakMap()): any {
  if (Object(obj) !== obj) return obj; // primitives
  if (hash.has(obj)) return hash.get(obj); // cyclic reference
  const result =
    obj instanceof Set
      ? new Set(obj) // See note about this!
      : obj instanceof Map
      ? new Map(Array.from(obj, ([key, val]) => [key, cloneDeep(val, hash)]))
      : obj instanceof Date
      ? new Date(obj)
      : obj instanceof RegExp
      ? new RegExp(obj.source, obj.flags)
      : // ... add here any specific treatment for other classes ...
      // and finally a catch-all:
      obj.constructor
      ? new obj.constructor()
      : Object.create(null);
  hash.set(obj, result);
  return Object.assign(
    result,
    ...Object.keys(obj).map(key => ({ [key]: cloneDeep(obj[key], hash) }))
  );
}

export const formateDate = (date: Date): string => {
  return `${date.toLocaleDateString()} ${date.toLocaleTimeString()}`;
};

export const normalizeDate = (dateStr?: string): string => {
  if (!dateStr) return new Date().toISOString().slice(0, -5);

  return new Date(dateStr).toISOString().slice(0, -5);
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
