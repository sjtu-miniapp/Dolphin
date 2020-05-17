import Taro, { FC, useState, useEffect } from '@tarojs/taro'
import { View, Swiper, SwiperItem } from '@tarojs/components';
import { BaseEventOrig } from '@tarojs/components/types/common';
import { SwiperProps } from '@tarojs/components/types/Swiper';
import { AtAvatar } from 'taro-ui';

import { Group, Task } from '../../types';
import * as groupAPI from '../../apis/groups';
import { getTasksByGroup } from '../../apis/tasks';
import FabButton from '../../components/fab-button';
import GroupModal from '../../components/group-modal';

import FullGroupView from './full-group';
import TaskView from './task';
import { GroupProps, ViewStatus } from './interface'

const GroupPage: FC<GroupProps> = _props => {

  const [viewStatus, setViewStatus] = useState<ViewStatus>('Full');

  const [groups, setGroups] = useState<Group[]>([]);

  const [showModal, setShowModal] = useState<boolean>(false);

  const updateGroups = async () => {
    try {
      const newGroups = await groupAPI.getGroupsByUser();
      setGroups(newGroups);
    } catch (error) {
      // TODO: remote data error handling
      console.error('Failed to update groups')
    }
  };

  useEffect(() => {
    updateGroups();
  }, []);

  const [selectedGroup, setSelectedGroup] = useState<Group | undefined>(undefined);

  const [tasks, setTasks] = useState<Task[]>([]);

  const updateTasks = async () => {
    try {
      if (!selectedGroup) return;
      const newTasks = await getTasksByGroup(selectedGroup.id);
      setTasks(newTasks);
    } catch (error) {
      // TODO: remote data error handling
      console.error('Failed to update groups')
    }
  };

  useEffect(() => {
    updateTasks();
  }, [selectedGroup]);

  const onSelectGroup = (groupID: string) => {
    const targetGroup = groups.find(g => g.id === groupID);

    if (!targetGroup) {
      console.error('Invalid selection group');
      return;
    }

    setViewStatus('Short');
    setSelectedGroup(targetGroup);
  }

  const onSelectTask = (taskID: string) => {
    console.log('todo: trigger task selection', taskID);
  }

  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);

  const addGroup = async (groupName: string) => {
    setShowModal(false);
    await groupAPI.addGroup(groupName);
    await updateGroups();
  }

  const handleSwipe = (e: BaseEventOrig<SwiperProps.onChangeEventDeatil>) => {
    showModal === true && closeModal();
    const { current } = e.detail;
    if (current === 0) setViewStatus('Full');
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
                onClickTask={onSelectTask}
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
