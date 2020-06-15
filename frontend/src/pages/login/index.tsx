import Taro, { FC, useState, useEffect, useDidShow } from '@tarojs/taro'
import { View } from '@tarojs/components';

import LoginHeader from '../../components/login-header';
import LoginFooter from '../../components/login-footer';

import * as auth from '../../apis/auth';

import './index.scss';

const LOGIN_EXPIRE_RANGE = 30 * 60 * 1000; // 30 mins

type UserInfo = auth.VerifyLoginData;

const Login: FC = () => {
  const [userInfo, setUserInfo] = useState<UserInfo | undefined>(undefined)
  const [isLogin, setIsLogin] = useState<boolean>(false);
  const [isLogout, setIsLogout] = useState<boolean>(false);

  const isLogged = !!userInfo;

  useDidShow(async () => {
    const lastLogin: number = Taro.getStorageSync('last_login');
    const userInfo: auth.VerifyLoginData = Taro.getStorageSync('userInfo');

    // Nothing was set in storage, probably haven't logged in before
    if (!userInfo || !lastLogin || isNaN(lastLogin)) {
      setUserInfo(undefined);
      return;
    }

    const { avatar, nickname, gender } = userInfo;
    if (!avatar || !nickname || !gender) {
      Taro.removeStorageSync('userInfo');
      setUserInfo(undefined);
      return;
    }

    // last_login is set in storage, but too far, automatically reload
    if (Date.now() - lastLogin >= LOGIN_EXPIRE_RANGE) {
      await handleLogin(avatar, nickname, gender);
    }

    // last_login is under 30 minutes
    else {
      setUserInfo({ avatar, nickname, gender });
    }
  })

  useEffect(() => {
    if (!userInfo) return;

    Taro.switchTab({ url: '/pages/group/index' });
  }, [userInfo]);

  const handleLogin = async (avatar: string, nickname: string, gender: 0 | 1 | 2) => {
    setIsLogin(true);

    const loginUserInfo = { avatar, nickname, gender };
    const loginResult = await auth.loginHandler(loginUserInfo);

    setIsLogin(false);
    if (!loginResult) {
      setUserInfo(undefined);
    } else {
      setUserInfo(loginUserInfo);
    }
  }

  const handleLogout = () => {
    setIsLogout(true);

    Taro.removeStorageSync('userInfo');
    Taro.removeStorageSync('last_login');

    setIsLogout(false);
    setUserInfo(undefined);
  }

  return (
    <View className='practice'>
      <LoginHeader
        isLogged={isLogged}
        isLogin={isLogin}
        userInfo={userInfo}
        onLogin={handleLogin}
      />
      <LoginFooter
        isLogged={isLogged}
        isLogout={isLogout}
        handleLogout={handleLogout}
      />
    </View>
  )
}

export default Login;
