import Taro, { FC, useState, useDidShow } from '@tarojs/taro'
import { View } from '@tarojs/components';
import bluebird from 'bluebird';

import { Group } from '../../types';
import * as groupAPI from '../../apis/group';
import * as shareAPI from '../../apis/share';
import FabButton from '../../components/fab-button';
import GroupModal, { InputType } from '../../components/group-modal';

import FullGroupView from './full-group';
import { normalizeGroup } from './utils';

const GroupPage: FC = () => {
  const [groups, setGroups] = useState<Group[]>([]);
  const [showModal, setShowModal] = useState<boolean>(false);

  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);

  const getGroupInfo = async () => {
    const groups = await groupAPI.getGroupsByUser();

    const groupDetails = await bluebird.map(
      groups,
      g => normalizeGroup(g)
    );

    return groupDetails.sort((prev, next) => next.updateTime.getTime() - prev.updateTime.getTime());
  }

  const updateGroups = async () => {
    try {
      const groupDetails = await getGroupInfo();
      console.log('Group Details:', JSON.stringify(groupDetails, null, 2));
      setGroups(groupDetails);
      Taro.setStorageSync('groups', groupDetails);
    } catch (error) {
      // TODO: remote data error handling
      console.error('Failed to update groups', error)
    }
  };

  useDidShow(updateGroups);

  const onSelectGroup = (groupID: number) => {
    const targetGroup = groups.find(g => g.id === groupID);

    if (!targetGroup) {
      console.error('Invalid selection group');
      return;
    }

    console.log(`TODO: on select group ${groupID}`);
    Taro.navigateTo({ url: `/pages/group-task/index?groupID=${groupID}` });
  }

  const onDeleteGroup = async (groupID: number) => {
    try {
      await groupAPI.deleteGroupByID(groupID);
      Taro.atMessage({ message: '成功删除小组!', type: 'success' });
    } catch (error) {
      Taro.atMessage({ message: '删除小组失败!', type: 'error' });
      console.error(error);
    }

    await updateGroups();
  }

  const addGroup = async (inputType: InputType, value: string) => {
    closeModal();
    if (inputType === 'groupName') {
      await groupAPI.createGroup({ name: value, user_ids: [] });
      await updateGroups();
    } else {
      const code = parseInt(value, 10);
      const result = await shareAPI.joinByCode(code);
      console.log('share joined !!!!!!', result);
    }
  }

  return (
    <View>
      <FullGroupView groups={groups} onClickGroup={onSelectGroup} onDeleteGroup={onDeleteGroup} />
      <FabButton onClick={openModal} />
      <GroupModal isOpened={showModal} handleCancel={closeModal} handleClose={closeModal} handleConfirm={addGroup} />
    </View>
  );
}

GroupPage.config = {
  navigationBarTitleText: '小组'
};

export default GroupPage;

