import Taro from "@tarojs/taro";

import config from "./config";
import * as utils from "./utils";
import { User } from "./group";

const HOST = config.BACKEND_HOST;
const BASE_URL = "/api/v1/task";
const PREFIX = `${HOST}${BASE_URL}`;

export interface TaskMeta {
  name: string;
  type: number;
  done: boolean;
  groupId: number;
  publisher_id: string;
  readonly: boolean;
  description: string;
}

export const getTasksByGroupID = async (
  groupID: number
): Promise<TaskMeta[]> => {
  const url = `${PREFIX}/${groupID}/group?${utils.getSessionQuery()}`;
  const response = await Taro.request<TaskMeta[]>({ url, method: "GET" });
  console.log(
    "Get Tasks By Group ID Result:",
    response.statusCode,
    response.data,
    response.errMsg
  );
  return response.data || [];
};

export type TaskWorker = User & { done: boolean };

export const getTaskWorker = async (taskID: number): Promise<TaskWorker[]> => {
  const url = `${PREFIX}/${taskID}/workers?${utils.getSessionQuery()}`;
  const response = await Taro.request<TaskWorker[]>({ url, method: "GET" });
  console.log(
    "Get Task Workers By Task ID Result:",
    response.statusCode,
    response.data,
    response.errMsg
  );

  if (response.statusCode !== 200) {
    console.error(
      `Failed to get task workers by id ${taskID}, on status: ${response.statusCode}, error: ${response.errMsg}`
    );
    return [];
  }

  return response.data;
};

interface UpdateTaskMeta {
  name: string;
  end_date: string;
  readonly: boolean;
  description: string;
  done: boolean;
}

export const updateTaskMeta = async (
  taskID: string,
  params: UpdateTaskMeta
): Promise<void> => {
  const url = `${PREFIX}/${taskID}/meta?${utils.getSessionQuery()}`;
  const response = await Taro.request<TaskWorker[]>({
    url,
    method: "POST",
    data: params
  });
  console.log(
    "Get Task Workers By Task ID Result:",
    response.statusCode,
    response.data,
    response.errMsg
  );
};

/*
		GroupId     int32      `json:"group_id"`
		UserIds     []string `json:"user_ids"`
		Name        string   `json:"name"`
		Type        int32      `json:"type"`
		LeaderId    *string   `json:"leader_id, omitempty"`
		StartDate   *string `json:"start_date, omitempty"`
		EndDate     *string `json:"end_date, omitempty"`
		Description *string   `json:"description, omitempty"`
		Readonly    bool     `json:"readonly"`
*/
export interface CreateTaskParams {
  group_id: number;
  user_ids: string[];
  name: string;
  type: number;
  end_date: string;
  description: string;
  readonly: boolean;
}

export const createTask = async (p: CreateTaskParams) => {
  const url = `${PREFIX}?${utils.getSessionQuery()}`;
  const response = await Taro.request<TaskWorker[]>({
    url,
    method: "PUT",
    data: p
  });
  console.log(
    "Create Task Result:",
    response.statusCode,
    response.data,
    response.errMsg
  );
};
