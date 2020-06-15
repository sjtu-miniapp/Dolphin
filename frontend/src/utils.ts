import moment from "moment";
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
  return moment(date).format("YYYY-MM-DD HH:mm:ss");
};

export const normalizeDate = (dateStr?: string): string => {
  if (!dateStr) return new Date().toISOString().slice(0, -5);

  return new Date(dateStr).toISOString().slice(0, -5);
};

export const normalizeTask = async (
  taskMeta: taskAPI.TaskMeta
): Promise<Task> => {
  const rawReceivers: taskAPI.TaskWorker[] = [];

  try {
    const receiversFromRemoteCall = await taskAPI.getTaskWorker(taskMeta.id);
    rawReceivers.push(...receiversFromRemoteCall);
  } catch (error) {
    console.error(`Failed to get task workers on: ${error}`);
  }

  const receivers = rawReceivers.map(r => normalizeTaskWorker(r));
  const publisher = receivers.find(r => r.id === taskMeta.publisher_id);

  return {
    id: taskMeta.id,
    groupID: taskMeta.group_id,
    name: taskMeta.name,
    description: taskMeta.description,
    startDate: convertOddStrToDate(taskMeta.start_date),
    endDate: convertOddStrToDate(taskMeta.end_date),
    readOnly: taskMeta.readonly,
    status: taskMeta.done ? "Done" : "Undone",
    receivers,
    type:
      taskMeta.type === 0 ? "个人" : taskMeta.type === 1 ? "团队" : undefined,
    publisher: publisher ? publisher.userName : ""
  };
};

export const convertOddStrToDate = (str: string): Date => {
  return moment.utc(str, "YYYY-M-D H:m:s").toDate();
};

export const normalizeTaskWorker = (
  taskWorker: taskAPI.TaskWorker
): UserTaskStatus => {
  return {
    id: taskWorker.id,
    userName: taskWorker.name,
    status: taskWorker.done ? "completed" : "in-progress"
  };
};
