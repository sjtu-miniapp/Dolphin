import Taro, { FC } from '@tarojs/taro'
import { AtDivider } from 'taro-ui';

interface TaskDivderProps {
  content?: string;
}

const TaskDivider: FC<TaskDivderProps> = props => {
  return (
    <AtDivider
      className='divider'
      lineColor='#F7F7F4'
      fontColor='lightsteelblue'
      fontSize={28}
      content={props.content || undefined}
    />
  )
}

export default TaskDivider;
