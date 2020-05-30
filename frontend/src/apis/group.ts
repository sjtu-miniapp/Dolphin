import Taro from "@tarojs/taro";

import config from "./config";
import * as utils from "./utils";

const HOST = config.BACKEND_HOST;
const BASE_URL = "/api/v1/group";

const createUrl = () => `${HOST}${BASE_URL}?${utils.getSessionQuery()}`;

export interface User {
  avatar: string;
  gender: number;
  id: string; // openid
  name: string;
  self_group_id: number;
}

interface Group {
  creator_id: string;
  name: string;
  type: number;
  users: User[];
}

export const getGroupByID = async (id: number): Promise<Group> => {
  const url = `${HOST}${BASE_URL}/${id}?${utils.getSessionQuery()}`;
  const response = await Taro.request({ url, method: "GET" });
  console.log(
    "Get Group By ID Result:",
    response.statusCode,
    response.data,
    response.errMsg
  );
  return response.data;
};

export interface GroupShort {
  creator_id: string;
  id: number;
  name: string;
}

export const getGroupsByUser = async (): Promise<GroupShort[]> => {
  const url = createUrl();
  const response = await Taro.request({ url, method: "GET" });
  console.log(response.statusCode, response.data);
  console.log(
    "Get Group By User Result:",
    response.statusCode,
    response.data,
    response.errMsg
  );
  return response.data;
};

export interface CreateGroupParams {
  name: string;
  user_ids: string[];
}

export const createGroup = async (params: CreateGroupParams): Promise<void> => {
  const url = createUrl();
  const response = await Taro.request({ url, method: "PUT", data: params });
  console.log(
    "Create Group Result:",
    response.statusCode,
    response.data,
    response.errMsg
  );
};
