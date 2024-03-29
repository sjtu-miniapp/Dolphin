import Taro, { FC, useState, useEffect, useRouter, useDidShow } from '@tarojs/taro'
import { View } from '@tarojs/components';
import { AtAvatar, AtButton } from 'taro-ui';
import bluebird from 'bluebird';

import { Group, Task } from '../../types';
import * as GroupAPI from '../../apis/group';
import * as taskAPI from '../../apis/task';
import * as shareAPI from '../../apis/share';
import TaskList from '../../components/task-list';
import FabButton from '../../components/fab-button';
import * as utils from '../../utils';

import TaskModal from './task-modal';

import './index.scss';

const GroupTaskPage: FC = () => {
  const groups: Group[] = Taro.getStorageSync('groups') || [];

  const [selectedGroupID, setSelectedGroupID] = useState<number | null>(() => {
    const { groupID } = useRouter().params;
    if (!isNaN(groupID as any)) return parseInt(groupID, 10);
    if (groups && groups.length > 0) return groups[0].id;
    return null;
  })

  const [tasks, setTasks] = useState<Task[]>([]);

  const [isAddTaskLayoutOpened, setIsAddTaskLayoutOpened] = useState<boolean>(false);
  const openLayOut = () => setIsAddTaskLayoutOpened(true);
  const closeLayOut = () => setIsAddTaskLayoutOpened(false);


  const [openShareCode, setOpenShareCode] = useState<boolean>(false);

  const [shareCode, setShareCode] = useState<string | number>('邀请码');

  const getGroupShareCode = async () => {
    if (!openShareCode || !selectedGroupID) return;
    setShareCode('loading...');
    const shareCode = await shareAPI.getShareCode('group', selectedGroupID)
    setShareCode(shareCode);
  }

  useEffect(() => {
    getGroupShareCode()
  }, [openShareCode])

  const getTaskDetails = async (): Promise<Task[]> => {
    if (!selectedGroupID) return [];

    const tasks = await taskAPI.getTasksByGroupID(selectedGroupID);
    const taskDetails = await bluebird.map(
      tasks,
      t => utils.normalizeTask(t)
    );

    return taskDetails.sort((prev, next) => next.endDate.getTime() - prev.endDate.getTime());
  }

  const getGroupDetail = async () => {
    if (!selectedGroupID) return;
    const detail = await GroupAPI.getGroupByID(selectedGroupID);
    console.log('deeeeeeeeeetail', detail);
  }

  const updateTasks = async () => {
    try {
      await getGroupDetail();
      const taskDetails = await getTaskDetails();
      console.log('Get Tasks Details:', JSON.stringify(taskDetails, null, 2));
      setTasks(taskDetails);
    } catch (error) {
      // TODO: remote data error handling
      console.error('Failed to update tasks', error)
    }
  };

  useDidShow(updateTasks);

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

  const addTask = async (params: taskAPI.CreateTaskParams) => {
    try {
      await taskAPI.createTask(params);
      closeLayOut();
      await updateTasks();
    } catch (error) {
      Taro.atMessage({ message: error, type: 'error' })
    }
  }

  const onClickFabbutton = () => {
    openLayOut();
  }

  const onDeleteTask = async id => {
    console.log('delete task', id);
    try {
      await taskAPI.deleteTask(id);
      await updateTasks();
    } catch (error) {
      Taro.atMessage({ message: error, type: 'error' })
    }
  }


  const switchShareCodeDisplay = () => {
    setOpenShareCode(!openShareCode);
  }

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
        <View className='grouptitle'>
          {selectedGroup ? selectedGroup.name : ''}
          <AtButton circle size='small' loading={shareCode === 'loading'} onClick={switchShareCodeDisplay}>{shareCode}</AtButton>
        </View>
        <TaskList tasks={tasks} onClickDelete={onDeleteTask} />
      </View>
      <FabButton onClick={onClickFabbutton} />
      <TaskModal groupID={selectedGroupID} isOpened={isAddTaskLayoutOpened} handleClose={closeLayOut} handleAdd={addTask} />
    </View>
  )
}

GroupTaskPage.config = {
  navigationBarTitleText: '小组详情'
};

export default GroupTaskPage;
