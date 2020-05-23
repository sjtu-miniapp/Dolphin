import Taro, { FC } from '@tarojs/taro'
import { View, Image } from '@tarojs/components'

import avatar from '../images/icon/dolphin.jpg';

import './logged-user.scss';

interface LoggedUserProps {
  userInfo: {
    avatar: string;
    nickName: string;
  }
}

const DEFAULT_PROPS = {
  avatar: '',
  nickName: '',
}

const LoggedUser: FC<LoggedUserProps> = props => {
  const { userInfo = DEFAULT_PROPS } = props;

  const onImageClick = () => {
    Taro.previewImage({
      urls: [userInfo.avatar],
    });
  };

  return (
    <View className="logged-user">
      <Image
        src={userInfo.avatar || avatar}
        className="user-avatar"
        onClick={onImageClick}
      />
      <View className="nickName">
        {userInfo.nickName || '游客'}
      </View>
    </View>
  )
}

LoggedUser.defaultProps = {
  userInfo: DEFAULT_PROPS
}

export default LoggedUser;
