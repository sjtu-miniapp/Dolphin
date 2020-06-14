import Taro, { FC, useState, useEffect } from '@tarojs/taro';
import { View } from '@tarojs/components';
import { AtActivityIndicator } from 'taro-ui';

import * as utils from '../../utils';

import { TaskStatus } from '.';
import './meta.scss';

export interface RawMeta {
  publisher: string;
  type: string;
  startDate: Date;
  endDate: Date;
}

export interface TaskMetaProps {
  taskStatus: TaskStatus<RawMeta>;
  onUpdateEndDate: (v: Date) => void;
}

const TaskMetaInfo: FC<TaskMetaProps> = props => {
  const { taskStatus } = props;

  const [startDate, setStartDate] = useState<string>('');
  const [endDate, setEndDate] = useState<string>('');

  useEffect(() => {
    if (taskStatus === 'Not Started' || taskStatus === 'Error' || taskStatus === 'Loading') {
      setStartDate('');
      setEndDate('');
    } else {
      const startDateStr = utils.formateDate(taskStatus.startDate);
      setStartDate(startDateStr);

      const endDateStr = utils.formateDate(taskStatus.endDate);
      setEndDate(endDateStr);
    }
  }, [taskStatus])

  return (
    <View className='taskmeta'>
      <View className='publisher'>
        {
          taskStatus === 'Not Started' ? 'Not Started' :
            taskStatus === 'Loading' ? <AtActivityIndicator /> :
              taskStatus === 'Error' ? 'Error' :
                <View className='tag'>
                  {taskStatus.publisher}
                </View>
        }
      </View>
      <View className='start'>
        {
          taskStatus === 'Not Started' ? 'Not Started' :
            taskStatus === 'Loading' ? <AtActivityIndicator /> :
              taskStatus === 'Error' ? 'Error' :
                <View className='tag'>
                  发布时间: {startDate}
                </View>
        }
      </View>
      <View className='type'>
        {
          taskStatus === 'Not Started' ? 'Not Started' :
            taskStatus === 'Loading' ? <AtActivityIndicator /> :
              taskStatus === 'Error' ? 'Error' :
                <View className='tag'>
                  类型: {taskStatus.type}
                </View>
        }
      </View>
      <View className='end'>
        {
          taskStatus === 'Not Started' ? 'Not Started' :
            taskStatus === 'Loading' ? <AtActivityIndicator /> :
              taskStatus === 'Error' ? 'Error' :
                <View className='tag' onClick={() => console.log('TODO: open date editor')}>
                  截止时间: {endDate}
                </View>
        }
      </View>
    </View>
  )

}

export default TaskMetaInfo;
