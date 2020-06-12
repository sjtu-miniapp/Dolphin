import Taro, { FC } from '@tarojs/taro';
import { View, Text } from '@tarojs/components'
import { AtList, AtListItem, AtSwipeAction } from 'taro-ui'
import { Task } from 'src/types';

const formateDate = (date: Date): string => {
  return `${date.toLocaleDateString()} ${date.toLocaleTimeString()}`;
}

interface TaskListProps {
  tasks: Task[],
  onClickDelete: (id: number) => Promise<void>;
}

const TaskList: FC<TaskListProps> = props => {
  const swipeActionOptions = [
    {
      text: '删除',
      style: {
        backgroundColor: '#6190E8'
      }
    }
  ];

  return (
    <View>
      <Text>Todo List</Text>
      <AtList>
        {
          props.tasks.map(item => {
            return (
              <AtSwipeAction
                key={item.id}
                onClick={() => props.onClickDelete(item.id)}
                options={swipeActionOptions}
              >
                <AtListItem
                  title={item.name}
                  note={formateDate(item.endDate)}
                  arrow='right'
                  iconInfo={{ size: 25, color: '#78A4FA', value: 'bookmark' }}
                  onClick={() => { Taro.navigateTo({ url: `/pages/task/index?id=${item.id}` }) }}
                />
              </AtSwipeAction>
            )
          })
        }
      </AtList>
    </View>
  )
}

TaskList.defaultProps = {
  tasks: []
}

export default TaskList;
