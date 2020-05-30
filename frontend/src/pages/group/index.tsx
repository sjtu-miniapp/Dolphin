import Taro, { FC, useState, useEffect } from '@tarojs/taro'
import { View, Swiper, SwiperItem } from '@tarojs/components';
import { BaseEventOrig } from '@tarojs/components/types/common';
import { SwiperProps } from '@tarojs/components/types/Swiper';
import { AtAvatar } from 'taro-ui';
import bluebird from 'bluebird';

import { Group, Task } from '../../types';
import * as groupAPI from '../../apis/groups';
import * as groupAPIMock from '../../apis/group';
import * as taskAPIMock from '../../apis/task';
import FabButton from '../../components/fab-button';
import GroupModal from '../../components/group-modal';

import FullGroupView from './full-group';
import TaskView from './task';
import { ViewStatus } from './interface'
import { normalizeGroup, normalizeTask } from './utils';

const GroupPage: FC = () => {

  const [viewStatus, setViewStatus] = useState<ViewStatus>('Full'); // 'Full/short' view for group page

  const [groups, setGroups] = useState<Group[]>([]);

  const [showModal, setShowModal] = useState<boolean>(false);

  const [selectedGroup, setSelectedGroup] = useState<Group | undefined>(undefined);

  const [tasks, setTasks] = useState<Task[]>([]);

  const getGroupInfo = async () => {
    const groups = await groupAPIMock.getGroupsByUser();

    const groupDetails = await bluebird.map(
      groups,
      g => normalizeGroup(g)
    );

    return groupDetails;
  }

  const updateGroups = async () => {
    try {
      const groupDetails = await getGroupInfo();
      console.log('Group Details:', JSON.stringify(groupDetails, null, 2));
      setGroups(groupDetails);
    } catch (error) {
      // TODO: remote data error handling
      console.error('Failed to update groups', error)
    }
  };

  useEffect(() => {
    updateGroups();
  }, []);

  const getTaskDetails = async (): Promise<Task[]> => {
    if (!selectedGroup) return [];

    const tasks = await taskAPIMock.getTasksByGroupID(selectedGroup.id);
    const taskDetails = await bluebird.map(
      tasks,
      (t, i) => normalizeTask(t, i)
    );

    return taskDetails;
  }

  // FIX ME: do we need this wrap function?
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
    updateTasks();
  }, [selectedGroup]);

  const onSelectGroup = (groupID: number) => {
    const targetGroup = groups.find(g => g.id === groupID);

    if (!targetGroup) {
      console.error('Invalid selection group');
      return;
    }

    setSelectedGroup(targetGroup);
    setViewStatus('Short');
  }

  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);

  const addGroup = async (groupName: string) => {
    closeModal();
    await groupAPI.addGroup(groupName);
    await groupAPIMock.createGroup({ name: groupName, user_ids: [] });
    await updateGroups();
  }

  const handleSwipe = (e: BaseEventOrig<SwiperProps.onChangeEventDeatil>) => {
    showModal === true && closeModal();
    if (e.detail.current === 0) setViewStatus('Full');
    else setViewStatus('Short');
  }

  return (
    <View>
      <Swiper
        current={viewStatus === 'Full' ? 0 : 1}
        style={{ width: '100vh', height: '100vh' }}
        onAnimationFinish={handleSwipe}
        skipHiddenItemLayout={true}
      >
        <SwiperItem>
          <FullGroupView groups={groups} onClickGroup={onSelectGroup} />
        </SwiperItem>
        <SwiperItem>
          <View className='at-row' >
            <View className='at-col-1' style={{ marginRight: '20px' }}>
              {groups.map(g => {
                const { name, id } = g;
                const style = name === (selectedGroup && selectedGroup.name) ? { backgroundColor: '#78A4FA' } : {};
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
                selectedGroupName={selectedGroup && selectedGroup.name}
              />
            </View>
          </View>
        </SwiperItem>
      </Swiper>
      {viewStatus === 'Full' && <FabButton onClick={openModal} />}
      <GroupModal isOpened={showModal} handleCancel={closeModal} handleClose={closeModal} handleConfirm={addGroup} />
    </View>
  )
}

GroupPage.config = {
  navigationBarTitleText: '小组'
};

export default GroupPage;
