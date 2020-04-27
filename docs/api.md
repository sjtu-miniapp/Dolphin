# 首页1（默认页）
1. 获取用户权限（token 2小时内有效）
2. 获取用户所属小组、创建者
3. 添加小组

# 首页2（日历页）
1. 获取任务名、任务相关人员、任务类型、状态——显示未完成任务
2. 获取任务发布者、任务内容、DDL——显示具体任务

# 小组页
1. 获取小组下所有任务名称、任务发布者、DDL——罗列小组内所有任务
2. 获取任务相关人员、任务状态——标识未完成任务
3. 添加任务

# 任务页
1. 获取任务名称、发布者、内容、DDL、相关人员、任务类型——显示具体任务及相关人员，确认Status修改权限
2. 修改任务属性（具体内容、相关人员、DDL、任务类型、任务状态）——获取修改记录
3. （获取用户信息，添加评论）

---

# Base URL: /api/v1/

# 首页1（默认页）
1. 获取用户权限（token 2小时内有效）

```
- Login
route: /login
method: POST
request data:
{
  wechatID: string;
  token: string;
}

response header:
{
  set-cookie: string;
}

response data:
{
  userName: string;
  userID: string;
}

// Fix me: not sure how we could login when dealing with wechat,
// we should probably update above spec later on.
```

2. 获取用户所属小组、创建者

```
- Get one group
route: /groups/:groupID
method: GET

request header:
{
  cookie: string;
}

response:
When succeed, return:
status: 200
data:
{
  group: { ... } // Whole data for one group
}

When failed, return:
status: 40x // Status code should align with the reason, e.g: 404 -> group not found
data:
{
  error: { ... } // Error message, since we can directly get indication from the status code, this field is optional.
}

- Get group by userID
route: /groups?userID=<user-id>
method: GET

request header:
{
  cookie: string;
}

response:
status: 200
data:
{
  groups: [
    {...},
    {...},
    {...},
    ...
  ]
}
```


3. 添加小组

```
- Create group
route: /groups
method: POST

request header:
{
  cookie: string;
}

request data:
{ ... } // Add necessary fields for creating a group, should not include id.

response:
status: 200
data:
{
  group: {...} // Whole group object created by request
}

- Update group
route: /groups/:groupID
method: POST

request header:
{
  cookie: string;
}

request data:
{ ... } // Add necessary fields for updating a group, should not include id.

response:
status: 200
data:
{
  group: {...} // Whole group object updated by request
}


- Delete group
route: /groups/:groupID
method: DELETE

request header:
{
  cookie: string;
}

response:
status: 200
data:
{
  group: {...} // Whole group object deleted by request
}
```


# 首页2（日历页）
1. 获取任务名、任务相关人员、任务类型、状态——显示未完成任务

```
- Get one task
route: /tasks/:taskID/short
method: GET

request header:
{
  cookie: string;
}

response:
data:
{
  task: { ... } // we should need the brief data (name, people, type, status) for one task
}

```

2. 获取任务发布者、任务内容、DDL——显示具体任务

```
- Get one task
route: /tasks/:taskID
method: GET

request header:
{
  cookie: string;
}

response:
data:
{
  task: { ... } // we should need the whole data for one task
}
```

# 小组页
1. 获取小组下所有任务名称、任务发布者、DDL——罗列小组内所有任务

```
- Get tasks by groupID
route: /tasks?groupID=<group-id>
method: GET

request header:
{
  cookie: string;
}

response:
data:
{
  tasks: [
    {...},
    {...},
    {...},
    ...
  ] 
}
```

2. 获取任务相关人员、任务状态——标识未完成任务

3. 添加任务

```
- Create task
route: /tasks
method: POST

request header:
{
  cookie: string;
}

request data:
{ ... } // Add necessary fields for creating a task.

response:
status: 200
data:
{
  task: {...} // Whole task object created by request
}
```


# 任务页
1. 获取任务名称、发布者、内容、DDL、相关人员、任务类型——显示具体任务及相关人员，确认Status修改权限

```
- Update a task
route: /tasks/:taskID
method: POST

request header:
{
  cookie: string;
}

request data:
{ ... } // Add necessary fields for updating a task, should not include id.

response:
status: 200
data:
{
  task: {...} // Whole task object updated by request
}


- Delete task
route: /tasks/:taskID
method: DELETE

request header:
{
  cookie: string;
}

response:
status: 200
data:
{
  task: {...} // Whole task object deleted by request
}
```

2. 修改任务属性（具体内容、相关人员、DDL、任务类型、任务状态）——获取修改记录

```
// TODO: Design form of data for tasks modification records (in frontend, we want to display sth. like `git-diff` or `color-diff` for each modification)
```

3. （获取用户信息，添加评论）

```
- Get comments for one task
route: /tasks/:taskID/comments
method: GET

request header:
{
  cookie: string;
}

response:
data:
{
  comments: [
    {...},
    {...},
    {...},
    ...
  ]
}

- Create task
route: /tasks/:taskID/comments
method: POST

request header:
{
  cookie: string;
}

request data:
{ ... } // Add necessary fields for creating a comment.

response:
status: 200
data:
{
  comment: {...} // Whole comment object created by request
}


- Update a comment
route: /tasks/:taskID/comment/:commentID
method: POST

request header:
{
  cookie: string;
}

request data:
{ ... } // Add necessary fields for updating a comment, should not include id.

response:
status: 200
data:
{
  comment: {...} // Whole comment object updated by request
}

- Delete comment
route: /tasks/:taskID/comment/:commentID
method: DELETE

request header:
{
  cookie: string;
}

response:
status: 200
data:
{
  comment: {...} // Whole comment object deleted by request
}
```
