import Taro, { FC } from '@tarojs/taro';
import { View } from '@tarojs/components';
import { AtMessage } from 'taro-ui';

import LoggedUser from './logged-user';
import LoginButton from './login-button';

import './login-header.scss';

interface LoginHeaderProps {
  userInfo: {
    avatar: string;
    nickName: string;
  };
  isLogged: boolean;
  setLoginInfo: (avatar: string, nickName: string) => void;
}

const LoginHeader: FC<LoginHeaderProps> = props => {
  return (
    <View className='user-box'>
      <AtMessage />
      <LoggedUser userInfo={props.userInfo} />
      {!props.isLogged && (
        <View className='login-button-box'>
          <LoginButton setLoginInfo={props.setLoginInfo} />
        </View>
      )}
    </View>
  )
}

LoginHeader.defaultProps = {
  userInfo: {
    avatar: '',
    nickName: '',
  },
  isLogged: false,
  setLoginInfo: () => { }
};

export default LoginHeader;
