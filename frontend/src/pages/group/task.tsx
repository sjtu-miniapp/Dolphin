import Taro, { FC } from '@tarojs/taro';
import { View } from '@tarojs/components';

import { TaskViewProps } from './interface'
import TaskList from '../../components/task-list';

const TaskView: FC<TaskViewProps> = props => {
  const { selectedGroupName, tasks } = props;

  return (
    <View>
      {selectedGroupName}
      <TaskList tasks={tasks} />
    </View>
  );
}

TaskView.defaultProps = {
  tasks: [],
  selectedGroupName: undefined,
}

export default TaskView;
