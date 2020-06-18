import Taro, { Component, Config } from '@tarojs/taro'
import { View } from '@tarojs/components'
import { AtCalendar } from 'taro-ui';

import TaskList from '../../components/task-list';
import { Task } from '../../types';
import './index.scss'

interface KanbanProps { }

interface KanbanState {
  list: Task[];
  inputVal: string;
}

export class Kanban extends Component<KanbanProps, KanbanState> {

  constructor(props: KanbanProps) {
    super(props);

    this.state = {
      list: [
        {
          id: 12,
          groupID: 123,
          name: "宏观经济学 期末作业",
          description: "",
          endDate: new Date("2020-06-30T00:00:00.000Z"),
          startDate: new Date(),
          readOnly: false,
          status: 'undone',
          receivers: [],
          type: '0'
        }
      ],
      inputVal: ''
    };
  }

  componentWillMount() { }

  componentDidMount() { }

  componentWillUnmount() { }

  componentDidShow() { }

  componentDidHide() { }

  /**
   * 指定config的类型声明为: Taro.Config
   *
   * 由于 typescript 对于 object 类型推导只能推出 Key 的基本类型
   * 对于像 navigationBarTextStyle: 'black' 这样的推导出的类型是 string
   * 提示和声明 navigationBarTextStyle: 'black' | 'white' 类型冲突, 需要显示声明类型
   */
  config: Config = {
    navigationBarTitleText: '首页'
  }

  render() {
    const { list } = this.state;

    return (
      <View className='index'>
        <AtCalendar currentDate='2020-06-30' />
        <TaskList tasks={list} onClickDelete={async () => { }} />
      </View>
    )
  }
}
