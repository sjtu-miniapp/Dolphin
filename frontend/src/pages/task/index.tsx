import Taro, { FC, useState, useEffect, useRouter, useDidShow } from '@tarojs/taro';
import { View } from '@tarojs/components';
import { AtTextarea/*, AtFloatLayout, AtCalendar*/ } from 'taro-ui';
import { BaseEventOrig } from '@tarojs/components/types/common';
import { InputProps } from '@tarojs/components/types/Input';
import { TagInfo } from 'taro-ui/types/tag';

import TaskReceiverList from '../../components/task-receiver-list';
import TaskDivider from '../../components/task-divider';
import * as taskAPI from '../../apis/task';
import { Task } from '../../types';
import * as utils from '../../utils';

import TaskTitle from './title';
import TaskMetaInfo, { RawMeta } from './meta';
import './index.scss';

type RemoteTask = Task | 'Not Started' | 'Loading' | 'Error';

export type TaskStatus<T> = 'Not Started' | 'Loading' | 'Error' | T;

const TaskView: FC = () => {
  const [taskDetail, setTaskDetail] = useState<RemoteTask>('Not Started');
  const [tempTask, setTempTask] = useState<RemoteTask>('Not Started');

  const taskID = useRouter().params.id;

  const updateTaskDetail = async (id: string): Promise<void> => {
    setTaskDetail('Loading');

    const rawTask = await taskAPI.getTaskMeta(id);
    const task = await utils.normalizeTask(rawTask);

    setTaskDetail(task || 'Error');
  };

  useDidShow(async () => {
    if (!taskID) {
      setTaskDetail('Not Started');
      return;
    }

    await updateTaskDetail(taskID);
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

  // const openEndTimeEdition = () => setOpenEndTimeEditing(true);
  // const closeEndTimeEdition = () => setOpenEndTimeEditing(false);

  // const onTaskEndDateChange = (item: { value: string }) => {
  //   const newTempTask = utils.cloneDeep(tempTask as Task);
  //   newTempTask.endDate = new Date(item.value);
  //   setTempTask(newTempTask);
  // }

  const onTaskUpdate = async () => {
    if (JSON.stringify(tempTask) === JSON.stringify(taskDetail)) return;

    if (tempTask === 'Not Started' || tempTask === 'Loading' || tempTask === 'Error') return;

    await taskAPI.updateTaskMeta(taskID, {
      name: tempTask.name,
      start_date: utils.normalizeDate(tempTask.startDate.toISOString()),
      end_date: utils.normalizeDate(tempTask.endDate.toISOString()),
      readonly: tempTask.readOnly,
      description: tempTask.description,
      done: tempTask.status === 'done'
    });

    await updateTaskDetail(taskID);
  }

  const titleStatus: TaskStatus<{ name: string; status: string }> =
    tempTask === 'Not Started' || tempTask === 'Loading' || tempTask === 'Error' ? tempTask
      : { name: tempTask.name, status: tempTask.status };

  const onUpdateStatus = async (v: TagInfo) => {
    console.log(JSON.stringify(v, null, 2));
  }

  const metaInfoStatus: TaskStatus<RawMeta> =
    tempTask === 'Not Started' || tempTask === 'Loading' || tempTask === 'Error' ? tempTask
      : { publisher: tempTask.publisher || '', type: tempTask.type || '', startDate: tempTask.startDate, endDate: tempTask.endDate };
  console.log('66666666666666', JSON.stringify(metaInfoStatus, null, 2));

  return (
    <View className='task'>
      <TaskTitle taskStatus={titleStatus} onUpdateName={onTaskNameChange} onBlurNameInput={onTaskUpdate} onUpdateStatus={onUpdateStatus} />
      <TaskMetaInfo taskStatus={metaInfoStatus} onUpdateEndDate={() => console.log('TODO: update end date')} />
      <TaskDivider content='任务详情' />
      <View className='description'>
        <AtTextarea
          height={500}
          count={false}
          maxLength={500}
          onChange={onTaskContentChange}
          onBlur={onTaskUpdate}
          value={tempTask === 'Not Started' || tempTask === 'Loading' || tempTask === 'Error' ? '' : tempTask.description}
        />
      </View>
      <TaskDivider content='任务进度' />
      <TaskReceiverList receivers={tempTask === 'Not Started' || tempTask === 'Loading' || tempTask === 'Error' ? [] : tempTask.receivers} />
      {/* <AtFloatLayout isOpened={openEndTimeEditing} onClose={closeEndTimeEdition}>
        <AtCalendar currentDate={tempTask.endDate} onDayClick={onTaskEndDateChange} />
      </AtFloatLayout> */}
    </View>
  )
}

TaskView.config = {
  navigationBarTitleText: '看板'
};

export default TaskView;
