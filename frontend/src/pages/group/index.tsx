import Taro, { Component, Config } from '@tarojs/taro'
import { View, Text, Navigator } from '@tarojs/components'
import './index.scss'

interface Group {
  name: string;
}

interface GroupProps { }

interface GroupState {
  groups: Group[];
}

export class Index extends Component<GroupProps, GroupState> {

  constructor(props: GroupProps) {
    super(props);

    this.state = {
      groups: [
        { name: '商业模式分析' },
        { name: '宏观经济学' },
        { name: '职业发展规划' },
      ],
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
    const { groups } = this.state;

    return (
      <View className='group'>
        <View className='group_list'>
          {
            groups.map(item => (
              <View className='item'>
                <Navigator url={'/'}>
                  {/* <Image className='img' src={item.picUrl}></Image> */}
                  <View className='right'>
                    <View className='text'>
                      <Text className='name'>{item.name}</Text>
                    </View>
                  </View>
                </Navigator>
              </View>
            ))
          }
        </View>
      </View>
    )
  }
}
