import Taro, { Component, Config } from '@tarojs/taro'
import { View, Text, Navigator, Image } from '@tarojs/components'
import './index.scss'

const GENERAL_URL = 'https://pic.downk.cc/item/5ea99a46c2a9a83be5804922.png';

interface Group {
  name: string;
  taskNumber: number;
  updateTime: Date;
  picUrl?: string;
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
        { name: '商业模式分析', taskNumber: 10, updateTime: new Date() },
        { name: '宏观经济学', taskNumber: 1, updateTime: new Date() },
        { name: '职业发展规划', taskNumber: 0, updateTime: new Date() },
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
                  <View className='left'>
                    <View className='text'>
                      <Text className='name'>{item.name}</Text>
                      <Text className='count'>任务数: {item.taskNumber}</Text>
                      <Text className='update'>上次更新: {item.updateTime.toLocaleDateString()}</Text>
                    </View>
                  </View>
                  <Image
                    className='right'
                    src={item.picUrl ? item.picUrl : GENERAL_URL}
                    mode='scaleToFill'
                    style={{ width: '132px', height: '99px', float: 'right', marginTop: '10px' }}
                  ></Image>
                </Navigator>
              </View>
            ))
          }
        </View>
      </View>
    )
  }
}
