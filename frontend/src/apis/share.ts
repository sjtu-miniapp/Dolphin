import Taro from "@tarojs/taro";

import config from "./config";
import * as utils from "./utils";

const HOST = config.BACKEND_HOST;
const BASE_URL = "/api/internal/v1";
const PREFIX = `${HOST}${BASE_URL}`;

export const getShareCode = async (
  type: string,
  id: number
): Promise<number> => {
  const url = `${PREFIX}/code?${utils.getSessionQuery()}&type=${type}&id=${id}`;
  const response = await Taro.request({ url, method: "GET" });
  console.log(
    "Get Share Code Result:",
    response.statusCode,
    response.data,
    response.errMsg
  );
  return response.data || 0;
};

export const joinByCode = async (code: number): Promise<any> {
  const url = `${PREFIX}/code/${code}?${utils.getSessionQuery()}`;
  const response = await Taro.request({ url, method: "PUT" });
  console.log(
    "Get Share Code Result:",
    response.statusCode,
    response.data,
    response.errMsg
  );
  return response.data;
}
