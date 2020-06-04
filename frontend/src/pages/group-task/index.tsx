import Taro, { FC, useState, useEffect, useRouter } from '@tarojs/taro'
import { View } from '@tarojs/components';
import { AtAvatar } from 'taro-ui';
import bluebird from 'bluebird';

import { Group, Task } from '../../types';
import * as taskAPI from '../../apis/task';

import TaskView from './task';
import { normalizeTask } from './utils';

const GroupTaskPage: FC = () => {
  const groups: Group[] = Taro.getStorageSync('groups') || [];
  console.log('gggggggg', groups);

  const [selectedGroupID, setSelectedGroupID] = useState<number | null>(() => {
    const { groupID } = useRouter().params;
    console.log('giiiiiiiiiid', groupID);
    if (!isNaN(groupID as any)) return parseInt(groupID, 10);
    if (groups && groups.length > 0) return groups[0].id;
    return null;
  })

  const [tasks, setTasks] = useState<Task[]>([]);

  const getTaskDetails = async (): Promise<Task[]> => {
    if (!selectedGroupID) return [];

    const tasks = await taskAPI.getTasksByGroupID(selectedGroupID);
    const taskDetails = await bluebird.map(
      tasks,
      (t, i) => normalizeTask(t, i)
    );

    return taskDetails;
  }

  const updateTasks = async () => {
    try {
      const taskDetails = await getTaskDetails();
      console.log('Get Tasks Details:', JSON.stringify(taskDetails, null, 2));
      setTasks(taskDetails);
    } catch (error) {
      // TODO: remote data error handling
      console.error('Failed to update groups')
    }
  };

  useEffect(() => {
    if (!selectedGroupID) {
      setTasks([]);
      return;
    }

    updateTasks();
  }, [selectedGroupID]);

  const onSelectGroup = (groupID: number) => {
    const targetGroup = groups.find(g => g.id === groupID);

    if (!targetGroup) {
      console.error('Invalid selection group');
      return;
    }

    setSelectedGroupID(targetGroup.id);
  }

  const selectedGroup = groups.find(g => g.id === selectedGroupID);

  return (
    <View className='at-row' >
      <View className='at-col-2'>
        {groups.map(g => {
          const { name, id } = g;
          const style = id === selectedGroupID ? { backgroundColor: '#78A4FA' } : {};
          return (
            <View onClick={() => onSelectGroup(id)} style={{ marginLeft: '8px', marginBottom: '20px' }}>
              <AtAvatar
                customStyle={style}
                circle
                text={name}
              />
            </View>
          );
        })}
      </View>
      <View className='at-col'>
        <TaskView
          tasks={tasks}
          selectedGroupName={selectedGroup ? selectedGroup.name : undefined}
        />
      </View>
    </View>
  )
}

GroupTaskPage.config = {
  navigationBarTitleText: '小组详情'
};

export default GroupTaskPage;
