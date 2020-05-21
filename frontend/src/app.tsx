import Taro, { Component, Config } from '@tarojs/taro';

import Index from './pages/group/index';
import './app.scss';

// 如果需要在 h5 环境中开启 React Devtools
// 取消以下注释：
// if (process.env.NODE_ENV !== 'production' && process.env.TARO_ENV === 'h5')  {
//   require('nerv-devtools')
// }

class App extends Component {

  componentDidMount() { }

  componentDidShow() { }

  componentDidHide() { }

  componentDidCatchError() { }

  /**
   * 指定config的类型声明为: Taro.Config
   *
   * 由于 typescript 对于 object 类型推导只能推出 Key 的基本类型
   * 对于像 navigationBarTextStyle: 'black' 这样的推导出的类型是 string
   * 提示和声明 navigationBarTextStyle: 'black' | 'white' 类型冲突, 需要显示声明类型
   */
  config: Config = {
    pages: [
      'pages/group/index',
      'pages/kanban/index',
      'pages/task/index',
    ],
    window: {
      backgroundTextStyle: 'light',
      navigationBarBackgroundColor: '#fff',
      navigationBarTitleText: 'Dolphin',
      navigationBarTextStyle: 'black'
    },
    tabBar: {
      list: [
        {
          pagePath: 'pages/group/index',
          text: '小组', //TODO: find a better name
          iconPath: './images/icon/group-unselected.png',
          selectedIconPath: './images/icon/group-selected.png',
        },
        {
          pagePath: 'pages/kanban/index',
          text: '看板',
          iconPath: './images/icon/kanban-unselected.png',
          selectedIconPath: './images/icon/kanban-selected.png'
        },
      ],
      color: '#ABBCDA',
      selectedColor: '#425DBA',
      backgroundColor: '#fff',
      borderStyle: 'white'
    }
  }

  // 在 App 类中的 render() 函数没有实际作用
  // 请勿修改此函数
  render() {
    return (
      <Index />
    )
  }
}

Taro.render(<App />, document.getElementById('app'))
