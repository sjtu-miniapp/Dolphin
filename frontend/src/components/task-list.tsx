import Taro, { FC } from '@tarojs/taro';
import { View, Text } from '@tarojs/components'
import { AtList, AtListItem } from 'taro-ui'
import { Task } from 'src/types';

const TaskList: FC<{ tasks: Task[] }> = props => {
  return (
    <View>
      <Text>Todo List</Text>
      <AtList>
        {
          props.tasks.map(item => {
            return (
              <AtListItem
                title={item.name}
                note={item.endDate.toLocaleDateString() + ' ' + item.endDate.toLocaleTimeString()}
                arrow='right'
                iconInfo={{ size: 25, color: '#78A4FA', value: 'bookmark' }}
                onClick={() => { Taro.navigateTo({ url: `/pages/task/index?id=${item.id}` }) }}
              />
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
