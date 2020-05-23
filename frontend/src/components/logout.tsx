import Taro, { FC } from '@tarojs/taro';
import { AtButton } from 'taro-ui';

interface LogoutProps {
  loading: boolean;
  handleLogout: () => void;
}

const Logout: FC<LogoutProps> = props => {
  return (
    <AtButton
      type='secondary'
      full
      loading={props.loading}
      onClick={props.handleLogout}
    >
      退出登录
    </AtButton>
  );
}

Logout.defaultProps = {
  loading: false,
  handleLogout: () => { }
}

export default Logout;
