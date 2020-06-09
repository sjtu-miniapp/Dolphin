import * as Taro from "@tarojs/taro";

export const getSessionQuery = (): string => {
  // FIX ME: how to handle session expired?
  const openID = Taro.getStorageSync("openid");
  const sessionID = Taro.getStorageSync("sid");
  return `openid=${openID}&sid=${sessionID}`;
};
