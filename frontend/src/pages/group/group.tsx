import Taro, { FC } from '@tarojs/taro';
import { View, Text, Navigator, Image } from '@tarojs/components';
import { AtAvatar } from 'taro-ui';

import { GroupViewProps } from './interface';
import './group.scss'

const GROUP_IAMGE_URL = 'https://pic.downk.cc/item/5ea99a46c2a9a83be5804922.png';

const GroupView: FC<GroupViewProps> = props => {
  const { groups, viewStatus, onClickGroup, seletectGroup } = props;
  if (viewStatus === 'Full' || !seletectGroup) {
    return (
      <View>
        {groups.map(item => (
          <View className='item' onClick={() => onClickGroup(item.id)} >
            <Navigator url={'/'}>
              <View className='left'>
                <View className='text'>
                  <Text className='name'>{item.name}</Text>
                  <Text className='count'>任务数: {item.taskNumber}</Text>
                  <Text className='update'>上次更新: {item.updateTime.toLocaleDateString()}</Text>
                </View>
              </View>
              <Image
                className='right'
                src={item.picUrl || GROUP_IAMGE_URL}
                mode='scaleToFill'
                style={{ width: '132px', height: '99px', float: 'right', marginTop: '10px' }}
              ></Image>
            </Navigator>
          </View>
        ))}
      </View>)
  }

  return (
    <View>
      {groups.map(g => {
        const { name, id } = g;
        const style = name === (seletectGroup && seletectGroup.name) ? { backgroundColor: '#78A4FA' } : {};

        return (
          <View onClick={() => onClickGroup(id)} style={{ marginLeft: '8px', marginBottom: '20px' }}>
            <AtAvatar
              customStyle={style}
              circle
              text={name}
            />
          </View>
        );
      })}
    </View>
  );
}

export default GroupView;

GroupView.defaultProps = {
  groups: [],
  seletectGroup: undefined,
  viewStatus: 'Full',
  onClickGroup: () => { }
}
