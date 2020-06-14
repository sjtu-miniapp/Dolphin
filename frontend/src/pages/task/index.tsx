import Taro, { FC, useState, useEffect, useRouter, useDidShow } from '@tarojs/taro';
import { View, Input } from '@tarojs/components';
import { AtTextarea, AtFloatLayout, AtCalendar, AtTag } from 'taro-ui';
import { BaseEventOrig } from '@tarojs/components/types/common';
import { InputProps } from '@tarojs/components/types/Input';

import TaskReceiverList from '../../components/task-receiver-list';
import TaskDivider from '../../components/task-divider';
import * as taskAPI from '../../apis/task';
import { normalizeTask } from '../group/utils';
import { Task } from '../../types';
import * as utils from '../../utils';
import './index.scss';

type RemoteTask = Task | 'Not Started' | 'Loading' | 'Error';

const TaskView: FC = () => {
  const [taskDetail, setTaskDetail] = useState<RemoteTask>('Not Started');
  const [tempTask, setTempTask] = useState<RemoteTask>('Not Started');
  const [openEndTimeEditing, setOpenEndTimeEditing] = useState<boolean>(false);

  const taskID = useRouter().params.id;

  const updateTask = async (id: string): Promise<void> => {
    setTaskDetail('Loading');

    const rawTask = await taskAPI.getTaskMeta(id);
    const task = await normalizeTask(rawTask);

    setTaskDetail(task || 'Error');
  };

  useDidShow(async () => {
    if (!taskID) {
      setTaskDetail('Error');
      return;
    }

    await updateTask(taskID);
  })

  useEffect(() => {
    setTempTask(utils.cloneDeep(taskDetail));
  }, [taskDetail]);


  const onTaskNameChange = (e: BaseEventOrig<InputProps.inputEventDetail>) => {
    const newTempTask = utils.cloneDeep(tempTask as Task);
    newTempTask.name = e.detail.value;
    setTempTask(newTempTask);
  }

  const onTaskContentChange = (v: string) => {
    const newTempTask = utils.cloneDeep(tempTask as Task);
    newTempTask.description = v;
    setTempTask(newTempTask);
  }

  if (tempTask === 'Not Started') return <View>{tempTask}</View>;
  if (tempTask === 'Loading') return <View>{tempTask}</View>;
  if (tempTask === 'Error') return <View>{tempTask}</View>;

  const openEndTimeEdition = () => setOpenEndTimeEditing(true);
  const closeEndTimeEdition = () => setOpenEndTimeEditing(false);

  const onTaskEndDateChange = (item: { value: string }) => {
    const newTempTask = utils.cloneDeep(tempTask as Task);
    newTempTask.endDate = new Date(item.value);
    setTempTask(newTempTask);
  }

  const onTaskUpdate = async () => {
    if (JSON.stringify(tempTask) === JSON.stringify(taskDetail)) return;

    await taskAPI.updateTaskMeta(taskID, {
      name: tempTask.name,
      start_date: utils.normalizeDate(tempTask.startDate.toISOString()),
      end_date: utils.normalizeDate(tempTask.endDate.toISOString()),
      readonly: tempTask.readOnly,
      description: tempTask.description,
      done: tempTask.status === 'done'
    });

    await updateTask(taskID);
  }

  return (
    <View className='task'>
      <View className='title'>
        <View className='taskname'>
          <Input name='taskname' onInput={onTaskNameChange} value={tempTask.name} onBlur={onTaskUpdate} />
        </View>
        <AtTag
          name='status'
          type='primary'
          size='normal'
          circle
          customStyle={{ color: '#78a4f4', borderColor: '#78a4f4', backgroundColor: '#fff' }}>
          {(taskDetail as Task).status === 'Undone' ? '未完成' : '完成'}
        </AtTag>
      </View>
      <View className='at-row at-row--wrap'>
        <View className='at-col at-col-4'><View className='tag'>发布人: {tempTask.publisher || 'N/A'}</View></View>
        <View className='at-col at-col-8'><View className='tag'>发布时间: {utils.formateDate(tempTask.startDate)}</View></View>
        <View className='at-col at-col-4'><View className='tag'>类型: {tempTask.type || 'N/A'}</View></View>
        <View className='at-col at-col-8'><View className='tag' onClick={openEndTimeEdition}>截止时间: {utils.formateDate(tempTask.endDate)}</View></View>
      </View>
      <TaskDivider content='任务详情' />
      <View className='description'>
        <AtTextarea height={500} count={false} maxLength={500} onChange={onTaskContentChange} value={tempTask.description} onBlur={onTaskUpdate} />
      </View>
      <TaskDivider content='任务进度' />
      <TaskReceiverList receivers={tempTask.receivers} />
      <AtFloatLayout isOpened={openEndTimeEditing} onClose={closeEndTimeEdition}>
        <AtCalendar currentDate={tempTask.endDate} onDayClick={onTaskEndDateChange} />
      </AtFloatLayout>
    </View>
  )
}

TaskView.config = {
  navigationBarTitleText: '看板'
};

export default TaskView;
