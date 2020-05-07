// import Taro, { useState, useEffect, useRouter, FC } from '@tarojs/taro'
import Taro, { FC, useState, useRouter } from '@tarojs/taro';
import { View, Text } from '@tarojs/components';
import { AtAvatar, AtList, AtListItem, AtFab } from 'taro-ui';

import './index.scss';

const GroupTaskPage: FC = () => {

  // FIX ME: load groups dynamically
  const [groups, _setGroups] = useState([
    { name: "商业模式分析", taskNumber: 10, updateTime: new Date() },
    { name: "宏观经济学", taskNumber: 1, updateTime: new Date() },
    { name: "职业发展规划", taskNumber: 0, updateTime: new Date() }
  ]);

  const [selectedGroup, setSelectedGroup] = useState(useRouter().params.group || '');

  const [groupTaskList, _setGroupTaskList] = useState([
    {
      name: 'Get Up',
      deadline: new Date()
    },
    {
      name: 'Coding',
      deadline: new Date()
    },
    {
      name: 'Gu Gu Gu',
      deadline: new Date()
    },
  ])

  const sidebar = groups.map(g => {
    const { name } = g;
    const setGroupFn = () => setSelectedGroup(name);
    const style = name === selectedGroup ? { backgroundColor: '#78A4FA' } : {};
    return (
      <View onClick={setGroupFn} style={{ marginLeft: '8px', marginBottom: '20px' }}>
        <AtAvatar
          customStyle={style}
          circle
          text={name}
        />
      </View>
    );
  });

  return (
    <View className='at-row'>
      <View className='at-col at-col-2 at-col--wrap'>
        {sidebar}
      </View>

      <View className='at-col'>
        {selectedGroup}
        <View className='list_wrap'>
          <Text>Todo List</Text>
          <AtList>
            {
              groupTaskList.map(item => {
                return (
                  <AtListItem
                    title={item.name}
                    note={item.deadline.toLocaleDateString() + ' ' + item.deadline.toLocaleTimeString()}
                    arrow='right'
                    iconInfo={{ size: 25, color: '#78A4FA', value: 'bookmark', }}
                  // iconInfo={{ size: 25, color: '#425D8A', value: 'bookmark', }}
                  >
                  </AtListItem>
                )
              })
            }
          </AtList>
        </View>

        <View className='button'>
          <AtFab>
            <Text className='at-fab__icon at-icon at-icon-add'></Text>
          </AtFab>
        </View>
      </View>
    </View>
  )

}

export default GroupTaskPage;
