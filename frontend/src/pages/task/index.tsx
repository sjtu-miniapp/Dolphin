import Taro, { FC, useState, useEffect, useRouter } from '@tarojs/taro';
import { View, Text } from '@tarojs/components';

import TaskReceiverList from '../../components/task-receiver-list';
import TaskDivider from '../../components/task-divider';
import { getTaskByID } from '../../apis/tasks';
import { Task } from '../../types';
import './index.scss';

type RemoteTask = Task | 'Not Started' | 'Loading' | 'Error';

const TaskView: FC = () => {
  const [taskDetail, setTaskDetail] = useState<RemoteTask>('Not Started');

  const router = useRouter();
  const taskID = router.params.id;

  const updateTask = async (id: string) => {
    const task = await getTaskByID(id);
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

  if (taskDetail === 'Not Started') return <View>{taskDetail}</View>;
  if (taskDetail === 'Loading') return <View>{taskDetail}</View>;
  if (taskDetail === 'Error') return <View>{taskDetail}</View>;

  return (
    <View className='task'>
      <View className='title'>{taskDetail.name}</View>
      <View className='at-row at-row--wrap'>
        <View className='at-col at-col-6'><View className='tag'>发布人: {taskDetail.publisher || 'N/A'}</View></View>
        <View className='at-col at-col-6'><View className='tag'>发布时间: {taskDetail.startDate.toLocaleDateString() || 'N/A'}</View></View>
        <View className='at-col at-col-6'><View className='tag'>截止时间: {taskDetail.endDate.toLocaleDateString() || 'N/A'}</View></View>
        <View className='at-col at-col-6'><View className='tag'>类型: {taskDetail.type || 'N/A'}</View></View>
      </View>
      <TaskDivider />
      <View className='description'>
        <Text>{taskDetail.description}</Text>
      </View>
      <TaskDivider content='任务进度' />
      <TaskReceiverList receivers={taskDetail.receivers} />
    </View>
  )
}

TaskView.config = {
  navigationBarTitleText: '看板'
};

export default TaskView;
