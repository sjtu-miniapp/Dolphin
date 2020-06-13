import Taro, { FC } from '@tarojs/taro';
import { View } from '@tarojs/components';
import { AtMessage, AtButton } from 'taro-ui';
import { ButtonProps } from '@tarojs/components/types/Button';
import { BaseEventOrig } from '@tarojs/components/types/common';

import * as auth from '../apis/auth';

import LoggedUser from './logged-user';

import './login-header.scss';

interface LoginHeaderProps {
  userInfo: auth.VerifyLoginData | undefined;
  isLogged: boolean;
  isLogin: boolean;
  onLogin: (avatar: string, nickName: string, gender: 0 | 1 | 2) => void;
}

const LoginHeader: FC<LoginHeaderProps> = props => {
  const onGetUserInfo = async (e: BaseEventOrig<ButtonProps.onGetUserInfoEventDetail>) => {
    const { avatarUrl, nickName, gender } = e.detail.userInfo;
    await props.onLogin(avatarUrl, nickName, gender);
  }

  return (
    <View className='user-box'>
      <AtMessage />
      <LoggedUser userInfo={props.userInfo} />
      {!props.isLogged && (
        <View className='login-button-box'>
          <AtButton
            openType='getUserInfo'
            onGetUserInfo={onGetUserInfo}
            type='primary'
            className='login-button'
            loading={props.isLogin}
          >微信登录</AtButton>
        </View>
      )}
    </View>
  )
}

LoginHeader.defaultProps = {
  isLogged: false,
  isLogin: false,
  onLogin: () => { }
};

export default LoginHeader;
