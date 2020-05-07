import Taro, { FC } from '@tarojs/taro';
import { View } from '@tarojs/components';

import { TaskViewProps } from './interface'
import TaskList from '../../components/task-list';
import FabButton from '../../components/fab-button';

const TaskView: FC<TaskViewProps> = props => {
  const { selectedGroupName, tasks } = props;

  return (
    <View>
      {selectedGroupName}
      <TaskList tasks={tasks} />
      <FabButton />
    </View>
  );
}

TaskView.defaultProps = {
  tasks: [],
  selectedGroupName: undefined,
  onClickTask: () => { }
}

export default TaskView;

