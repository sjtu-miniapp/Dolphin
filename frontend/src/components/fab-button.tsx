import Taro, { FC } from '@tarojs/taro';
import { View } from '@tarojs/components';
import { AtFab, AtIcon } from 'taro-ui';

import './fab-button.scss';

interface AtFabProps {
  onClick?: () => void;
}

const voidFunc = () => { };

const FabButton: FC<AtFabProps> = props => {
  return (
    <View className='floating'>
      <AtFab onClick={props.onClick || voidFunc} size='small'>
        <AtIcon value='add'></AtIcon>
      </AtFab>
    </View>
  )
}

FabButton.defaultProps = {
  onClick: voidFunc
}

export default FabButton;
