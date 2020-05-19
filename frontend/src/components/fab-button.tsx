import Taro, { FC } from '@tarojs/taro';
import { View } from '@tarojs/components'
import { AtFab, AtIcon } from 'taro-ui'
import './fab-button.scss';

interface AtFabProps {
  onClick?: () => void;
}

const FabButton: FC<AtFabProps> = props => {
  const onClick = props.onClick || (() => { });

  return (
    <View className='button'>
      <AtFab onClick={onClick} size='small'>
        <AtIcon className='at-fab__icon at-icon at-icon-add' value='add' ></AtIcon>
      </AtFab>
    </View>
  )
}

FabButton.defaultProps = {
  onClick: () => { }
}

export default FabButton;