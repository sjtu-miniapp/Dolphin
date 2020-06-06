import Taro, { FC } from '@tarojs/taro';
import { View, Text, Image } from '@tarojs/components';
import { AtSwipeAction } from 'taro-ui';
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
  onDeleteGroup: (groupID: number) => void;
}

const FullGroupView: FC<FullGroupViewProps> = props => {
  const { groups, onClickGroup } = props;

  const onClickDelete = async (groupID: number) => {
    await props.onDeleteGroup(groupID);
  }

  return (
    <View>
      {groups.map(g => (
        <View className='item'>
          <AtSwipeAction autoClose options={[
            {
              text: '删除',
              style: {
                backgroundColor: '#FF4949'
              }
            }
          ]} onClick={() => onClickDelete(g.id)}>
            <View className='left' onClick={() => onClickGroup(g.id)} >
              <View className='text'>
                <Text className='name'>{g.name}</Text>
                <Text className='count'>任务数: {g.taskNumber}</Text>
                <Text className='update'>上次更新: {g.updateTime ? g.updateTime.toLocaleDateString() : 'N/A'}</Text>
              </View>
            </View>
            <View className='right' style={{ backgroundColor: 'white' }}>
              <Image
                className='right'
                mode='scaleToFill'
                src={GROUP_IAMGE_URL}
                style={IAMGE_STYLE}
              />
            </View>
          </AtSwipeAction>
        </View>
      ))
      }
    </View >)
}

export default FullGroupView;

FullGroupView.defaultProps = {
  groups: [],
  onClickGroup: () => { }
}
