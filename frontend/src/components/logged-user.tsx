import Taro, { FC } from '@tarojs/taro'
import { View, Image } from '@tarojs/components'

import * as auth from '../apis/auth';

import avatar from '../images/icon/dolphin.jpg';

import './logged-user.scss';

interface LoggedUserProps {
  userInfo: auth.VerifyLoginData | undefined;
}

const DEFAULT_PROPS: auth.VerifyLoginData = {
  avatar: '',
  nickname: '',
  gender: 0
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
        {userInfo.nickname || '游客'}
      </View>
    </View>
  )
}

LoggedUser.defaultProps = {
  userInfo: undefined
}

export default LoggedUser;
