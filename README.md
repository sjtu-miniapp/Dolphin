# 需求规约草稿
## 点击小程序图标进入首页
// TODO: 首页概念图
- 点击小组管理（默认页），罗列用户所属小组
- 点击个人看版，显示当月日历，每日任务和完成情况

## 点击小组卡片，进入小组页面
//TODO：小组概念图
- 未完成的任务前有红点标记

## 点击任务，进入任务管理界面
//TODO：任务界面概念图
- 具体任务名称、内容、发布者、发布时间、截至日期、涉及成员、任务类型
- 评论区（评论者、评论内容、评论时间）
- 未完成名单

# API基本设计
// TODO: Redesign API, following is just a demo for querying data.
// We need to design all of the CRUD.
## 1. Visiting home page
### 1.1 getting groups
```
request:
  route: /api/v1/user/:userID/groups
  method: POST
  data:
    {
        token: string;  // user authentication ???
    }

response:
  [
      {
          id: string;   // group ID
          name: string; // group name
      },
      ...
  ]
```

### 1.2 getting tasks
```
request:
  route: /api/v1/user/:userID/tasks
  method: POST
  data:
    {
        token: string;  // user authentication ???
    }

response:
  [
      {
          id: string;   // task ID
          deadline: string;
          title: string; // title
          content: string;
          groupID: string;
          groupName: string;
      },
      ...
  ]
```

## 2. Visiting group page
### 2.1 listing tasks
```
request:
  route: /api/v1/user/:userID/tasks?groupID=<group-id>&detailed=false
  method: POST
  data:
    {
        token: string;  // user authentication ???
    }

response:
  [
      {
          id: string;   // task ID
          deadline: string;
          title: string; // title
      },
      ...
  ]
```

### 2.2 getting a task
```
request:
  route: /api/v1/user/:userID/task/:taskID
  method: POST
  data:
    {
        token: string;  // user authentication ???
    }

response:
  [
      {
          id: string;   // task ID
          deadline: string;
          title: string; // title
          content: string;
          groupID: string;
          groupName: string;
          comments: [
              {
                  id: string;
                  username: string;
                  content: string;
              },
              ...
          ]
      },
      ...
  ]
```

### 2.3 ?? getting comments
```
request:
  route: /api/v1/user/:userID/task/:taskID/comments
  method: POST
  data:
    {
        token: string;  // user authentication ???
    }

response:
  [
      {
          id: string;
          username: string;
          content: string;
      },
      ...
  ]
```

# 迭代计划

项目周期：2020.4.24-2020.5.21

## 迭代1：
ddl: 2020.5.5
- 前端：初步完成小程序三个页面，包括首页、小组页面、任务页面
       能够进行扫码访问
       微信账号登陆
       ？Jaccount账号登录
- 后端
    + 初步API设计并实现
    + 数据库设计
    + 单机部署
    + 单元测试

## 迭代2:
ddl:2020.5.12
- 前端：优化界面
- 后端
    + 性能优化
    + 集群部署
    + 压力测试
    + 集成测试
## 迭代3:
ddl:2020.5.19
+ 后端
    + 实现多人在线协同编辑
    + 系统测试
