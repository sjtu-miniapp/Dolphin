import Taro, { FC, useState } from '@tarojs/taro';
import { ButtonProps } from '@tarojs/components/types/Button';
import { BaseEventOrig } from '@tarojs/components/types/common';
import { AtButton } from 'taro-ui';

import './login-button.scss';

interface LoginButtonProps {
  setLoginInfo: (avatarUrl: string, nickName: string) => void;
}

const LoginButton: FC<LoginButtonProps> = props => {
  const [isLogin, setIsLogin] = useState<boolean>(false);

  const onGetUserInfo = async (e: BaseEventOrig<ButtonProps.onGetUserInfoEventDetail>) => {
    setIsLogin(true);

    const { avatarUrl, nickName } = e.detail.userInfo;

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
