import Taro, { FC } from '@tarojs/taro';
import { View } from '@tarojs/components';

import Logout from './logout';

import './login-footer.scss';

interface LoginFooterProps {
  isLogged: boolean;
  isLogout: boolean;
  handleLogout: () => void;
}

const LoginFooter: FC<LoginFooterProps> = props => {
  return (
    <View className='login-footer'>
      {props.isLogged && (
        <Logout loading={!props.isLogout} handleLogout={props.handleLogout} />
      )}
      <View className='tuture-motto'>
        {props.isLogged ? 'From Dolpin with Love ❤️' : '未登录'}
      </View>
    </View>
  )

}

LoginFooter.defaultProps = {
  isLogged: false,
  isLogout: false,
  handleLogout: () => { }
};

export default LoginFooter;
