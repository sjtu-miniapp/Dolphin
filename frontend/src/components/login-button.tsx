import Taro, { FC, useState } from '@tarojs/taro';
import { ButtonProps } from '@tarojs/components/types/Button';
import { BaseEventOrig } from '@tarojs/components/types/common';
import { AtButton } from 'taro-ui';

import * as auth from '../apis/auth';

import './login-button.scss';

interface LoginButtonProps {
  setLoginInfo: (avatarUrl: string, nickName: string) => void;
}

const loginHandler = async (code: string, userInfo: auth.VerifyLoginData): Promise<boolean> => {
  console.log(111111);
  const loginIds = await auth.getOpenIDByCode(code);
  console.log(222222, loginIds);
  const loginSucceed = await auth.verifyLoginByOpenID(loginIds, userInfo);
  console.log(333333, loginSucceed);

  if (loginSucceed) {
    Taro.setStorageSync('openid', loginIds.openid);
    Taro.setStorageSync('sid', loginIds.sid);
  }

  return loginSucceed;
}

const LoginButton: FC<LoginButtonProps> = props => {
  const [isLogin, setIsLogin] = useState<boolean>(false);

  const onGetUserInfo = async (e: BaseEventOrig<ButtonProps.onGetUserInfoEventDetail>) => {
    setIsLogin(true);

    const { code } = await Taro.login();
    console.log('login code:', code);

    const { avatarUrl, nickName, gender } = e.detail.userInfo;
    const loginResult = await loginHandler(code, { avatar: avatarUrl, nickname: nickName, gender });
    if (!loginResult) {
      Taro.atMessage({ message: '登录失败', type: 'error' });
      return;
    }

    Taro.atMessage({ message: '登录成功', type: 'success' });

    await props.setLoginInfo(avatarUrl, nickName);

    setIsLogin(false);
  }

  return (
    <AtButton
      openType='getUserInfo'
      onGetUserInfo={onGetUserInfo}
      type='primary'
      className='login-button'
      loading={isLogin}
    >微信登录</AtButton>
  )
}

export default LoginButton;
