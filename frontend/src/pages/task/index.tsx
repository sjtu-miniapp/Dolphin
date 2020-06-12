import Taro, { FC, useState, useEffect, useRouter } from '@tarojs/taro';
import { View } from '@tarojs/components';
import { AtTextarea, AtFloatLayout, AtCalendar } from 'taro-ui';
import { cloneDeep } from '../../utils';

import TaskReceiverList from '../../components/task-receiver-list';
import TaskDivider from '../../components/task-divider';
import * as taskAPI from '../../apis/task';
import { normalizeTask } from '../group/utils';
import { Task } from '../../types';
import './index.scss';

type RemoteTask = Task | 'Not Started' | 'Loading' | 'Error';

const TaskView: FC = () => {
  const [taskDetail, setTaskDetail] = useState<RemoteTask>('Not Started');
  const [tempTask, setTempTask] = useState<RemoteTask>('Not Started');
  const [openEndTimeEditing, setOpenEndTimeEditing] = useState<boolean>(false);

  const router = useRouter();
  const taskID = router.params.id;

  const updateTask = async (id: string) => {
    const rawTask = await taskAPI.getTaskMeta(id);
    const task = await normalizeTask(rawTask);
    setTaskDetail(task || 'Error');
  };

  useEffect(() => {
    if (!taskID) {
      setTaskDetail('Error');
      return;
    }

    setTaskDetail('Loading');

    updateTask(taskID);
  }, [taskID])

  useEffect(() => {
    setTempTask(cloneDeep(taskDetail));
  }, [taskDetail]);


  const onTaskNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newTempTask = cloneDeep(tempTask as Task);
    newTempTask.name = e.target.value;
    setTempTask(newTempTask);
  }

  const onTaskContentChange = (v: string) => {
    const newTempTask = cloneDeep(tempTask as Task);
    newTempTask.description = v;
    setTempTask(newTempTask);
  }

  if (tempTask === 'Not Started') return <View>{tempTask}</View>;
  if (tempTask === 'Loading') return <View>{tempTask}</View>;
  if (tempTask === 'Error') return <View>{tempTask}</View>;

  const openEndTimeEdition = () => setOpenEndTimeEditing(true);
  const closeEndTimeEdition = () => setOpenEndTimeEditing(false);

  const onTaskEndDateChange = (item: { value: string }) => {
    const newTempTask = cloneDeep(tempTask as Task);
    newTempTask.endDate = new Date(item.value);
    setTempTask(newTempTask);
  }

  return (
    <View className='task'>
      <View className='title'><input name='taskname' onChange={onTaskNameChange} value={tempTask.name} /></View>
      <View className='at-row at-row--wrap'>
        <View className='at-col at-col-6'><View className='tag'>发布人: {tempTask.publisher || 'N/A'}</View></View>
        <View className='at-col at-col-6'><View className='tag'>发布时间: {tempTask.startDate.toLocaleDateString() || 'N/A'}</View></View>
        <View className='at-col at-col-6'><View className='tag' onClick={openEndTimeEdition}>截止时间: {tempTask.endDate.toLocaleDateString() || 'N/A'}</View></View>
        <View className='at-col at-col-6'><View className='tag'>类型: {tempTask.type || 'N/A'}</View></View>
      </View>
      <TaskDivider content='任务详情' />
      <View className='description'>
        <AtTextarea height={500} count={false} maxLength={500} onChange={onTaskContentChange} value={tempTask.description} />
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
