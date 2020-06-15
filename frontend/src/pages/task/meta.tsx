import Taro, { FC, useState, useEffect } from '@tarojs/taro';
import { View, Text, Picker } from '@tarojs/components';
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
  onUpdateEndDate: (v: Date) => Promise<void>;
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

  const [ddlDate = '', ddlTime = ''] = endDate.split(' ');

  const updateEndDate = async (replacePattern: { [Symbol.replace](string: string, replaceValue: string): string; }, replaceValue: string) => {
    if (taskStatus === 'Error' || taskStatus === 'Loading' || taskStatus === 'Not Started') return;

    const newDateStr = utils.formateDate(taskStatus.endDate).replace(replacePattern, replaceValue);
    const newEndDate = new Date(newDateStr);

    await props.onUpdateEndDate(newEndDate);
  }

  const onDateChange = async e => {
    await updateEndDate(/\d{4}-\d{2}-\d{2}/, e.detail.value)
  }

  const onTimeChange = async e => {
    await updateEndDate(/\d{2}:\d{2}/, e.detail.value);
  }

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
                  <Text space='ensp'>发布时间: </Text>
                  <View>{startDate}</View>
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
                <View className='tag'>
                  <Text space='ensp'>截止时间: </Text>
                  <Picker mode='date' onChange={onDateChange} value={ddlDate}>
                    <View >{ddlDate}</View>
                  </Picker>
                  <View className='split'>|</View>
                  <Picker mode='time' onChange={onTimeChange} value={ddlTime}>
                    <View>{ddlTime}</View>
                  </Picker>
                </View>
        }
      </View>
    </View>
  )
}

export default TaskMetaInfo;
