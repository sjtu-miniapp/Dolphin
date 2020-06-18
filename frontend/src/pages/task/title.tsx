import Taro, { FC } from '@tarojs/taro';
import { View, Input } from '@tarojs/components';
import { AtActivityIndicator, AtInput, AtTag } from 'taro-ui';
import { BaseEventOrig } from '@tarojs/components/types/common';
import { InputProps } from '@tarojs/components/types/Input';

import { TaskStatus } from '.';

import './title.scss';

export interface TaskTitleProps {
  taskStatus: TaskStatus<{
    name: string;
    status: string;
  }>;
  onUpdateName: (e: BaseEventOrig<InputProps>) => void;
  onBlurNameInput: () => void;
  onUpdateStatus: (value: 'done' | 'undone') => Promise<void>; // send updated value
}

const TaskTitle: FC<TaskTitleProps> = props => {
  const { taskStatus } = props;

  const onClickStatus = async () => {
    if (taskStatus === 'Error' || taskStatus === 'Loading' || taskStatus === 'Not Started') return;

    const value = taskStatus.status === 'done' ? 'undone' : 'done';
    await props.onUpdateStatus(value);
  }

  return (
    <View className='title'>
      <View className='taskname'>
        {
          taskStatus === 'Not Started' ? (<Input name='taskname' disabled />) :
            taskStatus === 'Loading' ? (<AtActivityIndicator />) :
              taskStatus === 'Error' ? (<AtInput error title='无法获取task' type='text' name='titleInput' onChange={() => { }} />) :
                <Input name='taskname' onInput={props.onUpdateName} value={taskStatus.name} onBlur={props.onBlurNameInput} />
        }
      </View>
      {
        taskStatus === 'Not Started' ? (<AtTag name='status' type='primary' circle />) :
          taskStatus === 'Loading' ? (<AtActivityIndicator />) :
            taskStatus === 'Error' ? (<View></View>) :
              taskStatus.status === 'Undone' ? (<AtTag name='status' circle onClick={onClickStatus} customStyle={{ color: '#78a4f4', borderColor: '#78a4f4', backgroundColor: '#fff' }}>未完成</AtTag>) :
                <AtTag name='status' circle customStyle={{ color: '#78a4f4', borderColor: '#78a4f4', backgroundColor: '#fff' }}>完成</AtTag>
      }
    </View>
  )
}

TaskTitle.defaultProps = {
  taskStatus: 'Not Started',
  onUpdateName: () => { },
  onBlurNameInput: () => { },
  onUpdateStatus: async () => { },
}

export default TaskTitle;
