import Taro, { FC, useState, useEffect } from '@tarojs/taro'
import { View } from '@tarojs/components';

import LoginHeader from '../../components/login-header';
import LoginFooter from '../../components/login-footer';

import './index.scss';

interface UserInfo {
  avatar: string;
  nickName: string;
}

const DEFAULT_USER: UserInfo = {
  avatar: '',
  nickName: ''
}

const Login: FC = () => {
  const [userInfo, setUserInfo] = useState<UserInfo | undefined>(undefined)
  const [isLogout, setIsLogout] = useState(false);

  const isLogged = !!userInfo;

  const setUserInfoFromStorage = async () => {
    try {
      const { data } = await Taro.getStorage({ key: 'userInfo' });
      if (!data || !data.nickName) return;

      const { nickName, avatar } = data;
      setUserInfo({ nickName, avatar });
    } catch (error) {
      console.error('getStorage ERR:', error);

    }
  }

  useEffect(() => { setUserInfoFromStorage() }, []);

  useEffect(() => {
    if (!userInfo || !userInfo.nickName) return;
    console.log(11111111, userInfo);
    Taro.switchTab({ url: '/pages/group/index' });
  }, [userInfo]);

  const setLoginInfo = async (avatar: string, nickName: string) => {
    try {
      await Taro.setStorage({
        key: 'userInfo',
        data: { avatar, nickName }
      });

      await setUserInfoFromStorage();

    } catch (error) {
      console.error('setStorage ERR:', error);
    }
  }

  const handleLogout = async () => {
    setIsLogout(true);

    try {
      await Taro.removeStorage({ key: 'userInfo' });
      setUserInfo(undefined);
    } catch (error) {
      console.error('removeStorage ERR:', error);
    }
  }

  return (
    <View className='practice'>
      <LoginHeader
        isLogged={isLogged}
        userInfo={userInfo || DEFAULT_USER}
        setLoginInfo={setLoginInfo}
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
