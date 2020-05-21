import Taro, { FC } from '@tarojs/taro'
import { View, Image } from '@tarojs/components'

import avatar from '../images/icon/kanban-selected.png';

interface LoggedUserProps {
  userInfo: {
    avatar: string;
    nickName: string;
    userName: string;
  }
}

const DEFAULT_PROPS = {
  avatar: '',
  nickName: '',
  userName: ''
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
      <View className="username">{userInfo.userName}</View>
    </View>
  )
}

LoggedUser.defaultProps = {
  userInfo: DEFAULT_PROPS
}

export default LoggedUser;
