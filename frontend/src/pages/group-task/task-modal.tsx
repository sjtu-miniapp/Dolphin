import Taro, { FC, useState, useEffect } from '@tarojs/taro'
import { View } from '@tarojs/components';
import { AtFloatLayout, AtInput, AtSwitch, AtAccordion, AtCalendar } from 'taro-ui';
import moment from 'moment';

// import TaskDivider from '../../components/task-divider';
import './task-modal.scss';

interface TaskModalProps {
  isOpened: boolean;
  handleClose: () => void;
  handleAdd: () => void;
}

export type TaskType = '个人' | '团队';

const TaskModal: FC<TaskModalProps> = props => {
  const handleClose = () => {
    props.handleClose()
  }

  const [taskName, setTaskName] = useState<string>('');
  const [taskType, setTaskType] = useState<TaskType>('个人');

  const [openCalendar, setOpenCalendar] = useState<boolean>(false);
  const [taskDeadlineDate, setTaskDeadlineDate] = useState<string>(moment().format('YYYY/MM/DD'));
  // const [taskDeadlineTime, setTaskDeadlineTime] = useState<string>(moment().format('hh:mm'));

  const onChangeTaskName = v => setTaskName(v);

  const onChangeTaskType = v => {
    if (v) setTaskType('团队');
    else setTaskType('个人');
  };

  const onChangeTaskDeadlineDate = v => setTaskDeadlineDate(v);
  // const onChangeTaskDeadlineTime = v => setTaskDeadlineTime(v);

  const switchCalenderOpen = () => {
    if (openCalendar) setOpenCalendar(false);
    else setOpenCalendar(true);
  }

  const onSelectDeadlinDate = (v: { value: { end: string; start: string; } }) => {
    setTaskDeadlineDate(v.value.end);
  }

  useEffect(() => {
    console.log(taskDeadlineDate);
  }, [taskDeadlineDate])

  return (
    <AtFloatLayout isOpened={props.isOpened} title='创建小组' scrollY onClose={handleClose} >
      <View className='task'>
        <AtInput
          name='taskname'
          type="text"
          title='任务名'
          required
          border
          value={taskName}
          onChange={onChangeTaskName}
        />
        <AtSwitch title={`类型：${taskType}`} checked={taskType === '团队'} onChange={onChangeTaskType} />
        <AtAccordion
          open={openCalendar}
          onClick={switchCalenderOpen}
          title={`截止日期：${taskDeadlineDate}`}
          icon={{ value: 'calendar', color: '#7FA6C9', size: '15' }}>
          <AtCalendar currentDate={taskDeadlineDate} onSelectDate={onSelectDeadlinDate} />
        </AtAccordion>
      </View>
    </AtFloatLayout>
  )
}

TaskModal.defaultProps = {
  isOpened: false,
  handleClose: () => { },
  handleAdd: () => { }
}

export default TaskModal;
