import Taro, { FC } from '@tarojs/taro';
import { AtTimeline } from 'taro-ui'
import { Item } from 'taro-ui/types/timeline'

import { UserTaskStatus } from '../types';

const TaskReceiverList: FC<{ receivers: UserTaskStatus[] }> = props => {
  const completed = props.receivers.filter(r => r.status === 'completed').map(r => r.userName).join(' ');
  const inProgess = props.receivers.filter(r => r.status === 'in-progress').map(r => r.userName).join(' ');
  const toStart = props.receivers.filter(r => r.status === 'to-start').map(r => r.userName).join(' ');

  const items: Item[] = [
    { title: '已完成', content: [completed], color: 'green', icon: 'check-circle' },
    { title: '进行中', content: [inProgess], color: 'yellow', icon: 'loading-2' },
    { title: '未开始', content: [toStart], color: 'red', icon: 'stop' },
  ]
  return <AtTimeline items={items} />;
}

TaskReceiverList.defaultProps = {
  receivers: []
}

export default TaskReceiverList;

