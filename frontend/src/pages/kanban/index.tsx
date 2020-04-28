import Taro, { Component, Config } from '@tarojs/taro'
import { View, Text, Input } from '@tarojs/components'
import './index.scss'

interface KanbanProps { }

interface KanbanState {
  list: string[];
  inputVal: string;
}

export class Kanban extends Component<KanbanProps, KanbanState> {

  constructor(props: KanbanProps) {
    super(props);

    this.state = {
      list: [
        'Get Up',
        'Coding',
        'Gu Gu Gu'
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

  inputHandler = e => this.setState({ inputVal: e.target.value });

  addTodo = () => {
    const { list, inputVal } = this.state;

    if (inputVal === '') return;

    this.setState({
      list: list.concat(inputVal), inputVal: ''
    });
  }

  deleteTodo = e => {
    const idx = parseInt(e.target.value, 10);

    const list = [...this.state.list];
    list.splice(idx, 1)

    this.setState({ list });

  }

  render() {
    const { list, inputVal } = this.state;

    return (
      <View className='index'>
        <Input className='input' type='text' value={inputVal} onInput={this.inputHandler} />
        <Text className='add' onClick={this.addTodo}>添加</Text>
        <View className='list_wrap'>
          <Text>Todo List</Text>
          {
            list.map((item, index) => {
              return (
                <View>
                  <Text>{index + 1}.{item}</Text>
                  <Text className='del' id={index.toString()} onClick={this.deleteTodo}>删除</Text>
                </View>
              )
            })
          }
        </View>
      </View>
    )
  }
}
