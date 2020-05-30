import Taro from "@tarojs/taro";

import config from "./config";

const HOST = config.BACKEND_HOST;
const BASE_URL = "/api/v1/auth";
const ON_LOGIN_ROUTE = "/on_login";
const VERIFY_LOGIN_ROUTE = "/after_login";

interface OnLogin {
  openid: string;
  sid: string;
}

export const getOpenIDByCode = async (code: string): Promise<OnLogin> => {
  const url = `${HOST}${BASE_URL}${ON_LOGIN_ROUTE}?code=${code}`;
  const response = await Taro.request({ url, method: "POST" });
  return response.data;
};

export interface VerifyLoginData {
  avatar: string;
  gender: number;
  nickname: string;
}

export const verifyLoginByOpenID = async (
  loginParams: OnLogin,
  userData: VerifyLoginData
): Promise<boolean> => {
  const url = `${HOST}${BASE_URL}${VERIFY_LOGIN_ROUTE}?openid=${loginParams.openid}&sid=${loginParams.sid}`;
  const response = await Taro.request({ url, method: "PUT", data: userData });
  if (response.statusCode >= 200 && response.statusCode < 300) return true;
  return false;
};
