import Taro, { FC } from '@tarojs/taro';
import { View } from '@tarojs/components'
import { AtFab, AtIcon } from 'taro-ui'
import './fab-button.scss';

const FabButton: FC = () => {
  return (
    <View className='button'>
      <AtFab onClick={() => { }} size='small'>
        <AtIcon className='at-fab__icon at-icon at-icon-add' value='add' ></AtIcon>
      </AtFab>
    </View>
  )
}

FabButton.defaultProps = {
  tasks: []
}

export default FabButton;
