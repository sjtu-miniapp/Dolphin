import Taro, { FC, useState, useEffect } from '@tarojs/taro'
import { View, ScrollView } from '@tarojs/components'

import { Group, Task } from 'src/types';
import * as groupAPI from '../../apis/groups';
import { getTasksByGroup } from '../../apis/tasks';

import GroupView from './group';
import TaskView from './task';
import { GroupProps, ViewStatus } from './interface'
import FabButton from '../../components/fab-button';
import GroupModal from '../../components/group-modal';

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

  const groupViewClassName = viewStatus === 'Full' ? 'at-col at-col-12' : 'at-col at-col-2';

  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);

  const addGroup = async (groupName: string) => {
    setShowModal(false);
    await groupAPI.addGroup(groupName);
    await updateGroups();
  }

  return (
    <View className='at-row'>
      <View className={groupViewClassName}>
        <GroupView groups={groups} viewStatus={viewStatus} onClickGroup={onSelectGroup} seletectGroup={selectedGroup} />
        <FabButton onClick={openModal} />
        <GroupModal isOpened={showModal} handleCancel={closeModal} handleClose={closeModal} handleConfirm={addGroup} />
      </View>
      <ScrollView
        style={{ whiteSpace: 'nowrap' }}
        scrollX
        scrollLeft={0}
        scrollWithAnimation
        onScrollToLower={() => setViewStatus('Full')}
      >
        {viewStatus === 'Full'
          ? <View />
          : <TaskView
            tasks={tasks}
            onClickTask={onSelectTask}
            selectedGroupName={selectedGroup && selectedGroup.name}
          />
        }
      </ScrollView>
    </View>
  )
}

GroupPage.config = {
  navigationBarTitleText: '小组'
};

export default GroupPage;
