import Taro, { FC } from '@tarojs/taro';
import { View, Text, Navigator, Image } from '@tarojs/components';
import { CSSProperties } from 'react';

import { Group } from '../../types';
import './group.scss'

const GROUP_IAMGE_URL = 'https://pic.downk.cc/item/5ea99a46c2a9a83be5804922.png';

const IAMGE_STYLE: CSSProperties = {
  width: '132px',
  height: '99px',
  float: 'right',
  marginTop: '10px'
}

interface FullGroupViewProps {
  groups: Group[];
  onClickGroup: (groupID: number) => void;
}

const FullGroupView: FC<FullGroupViewProps> = props => {
  const { groups, onClickGroup } = props;
  return (
    <View>
      {groups.map(g => (
        <View className='item' onClick={() => onClickGroup(g.id)} >
          <Navigator url={'/'}>
            <View className='left'>
              <View className='text'>
                <Text className='name'>{g.name}</Text>
                <Text className='count'>任务数: {g.taskNumber}</Text>
                <Text className='update'>上次更新: {g.updateTime ? g.updateTime.toLocaleDateString() : 'N/A'}</Text>
              </View>
            </View>
            <Image
              className='right'
              mode='scaleToFill'
              src={GROUP_IAMGE_URL}
              style={IAMGE_STYLE}
            ></Image>
          </Navigator>
        </View>
      ))}
    </View>)
}

export default FullGroupView;

FullGroupView.defaultProps = {
  groups: [],
  onClickGroup: () => { }
}
