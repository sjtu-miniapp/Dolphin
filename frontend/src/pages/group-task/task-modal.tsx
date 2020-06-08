import Taro, { FC, useState, useEffect } from '@tarojs/taro'
import { View } from '@tarojs/components';
import { AtFloatLayout, AtInput, AtSwitch, AtAccordion, AtCalendar, AtTextarea, AtButton } from 'taro-ui';
import moment from 'moment';

// import TaskDivider from '../../components/task-divider';
import { CreateTaskParams } from '../../apis/task';
import * as taskAPI from '../../apis/task';
import './task-modal.scss';

interface TaskModalProps {
  isOpened: boolean;
  handleClose: () => void;
  handleAdd: () => void;
  groupID: number | null;
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

  const [taskContent, setTaskContent] = useState<string>('');

  const onChangeTaskName = v => setTaskName(v);

  const onChangeTaskType = v => {
    if (v) setTaskType('团队');
    else setTaskType('个人');
  };

  const switchCalenderOpen = () => {
    if (openCalendar) setOpenCalendar(false);
    else setOpenCalendar(true);
  }

  const onSelectDeadlineDate = (v: { value: { end: string; start: string; } }) => {
    setTaskDeadlineDate(v.value.end);
  }

  const onTaskContentChange = v => setTaskContent(v);

  useEffect(() => {
    console.log(taskName, taskType, taskDeadlineDate, taskContent);
  }, [taskName, taskType, taskDeadlineDate, taskContent])

  const onClickCreateButton = async () => {
    if (!props.groupID) return;

    const user_ids = [Taro.getStorageSync('openid')];

    const param: CreateTaskParams = {
      group_id: props.groupID,
      user_ids,
      name: taskName,
      type: taskType === '个人' ? 0 : 1,
      end_date: taskDeadlineDate,
      description: taskContent,
      readonly: false
    };
    await taskAPI.createTask(param);
  }

  return (
    <AtFloatLayout isOpened={props.isOpened} title='创建任务' scrollY onClose={handleClose} >
      <View className='task'>
        <AtInput
          name='taskname'
          type="text"
          title='任务名:'
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
          <AtCalendar currentDate={taskDeadlineDate} onSelectDate={onSelectDeadlineDate} />
        </AtAccordion>
        {/* <TaskDivider content='任务详情' /> */}
        <View className='description'>
          <AtTextarea placeholder='任务详情......' height={500} count={false} maxLength={500} onChange={onTaskContentChange} value={taskContent} />
        </View>
        <View className='taskbutton'>
          <AtButton circle size='small' type='primary' onClick={onClickCreateButton}>提交</AtButton>
        </View>
      </View>
    </AtFloatLayout>
  )
}

TaskModal.defaultProps = {
  isOpened: false,
  handleClose: () => { },
  handleAdd: () => { },
  groupID: null
}

export default TaskModal;
