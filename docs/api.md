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
# onLogin: acquire openid and sessionid and then put them in storage
- route: /auth/on_login
- method: POST
- request data:
  - code string
- response data:
  - openid string
  - sid string
- response status:
  - 200 success
  - 500 failure

# afterLogin: callback of onLogin; get userInfo through wx api
- route: /auth/after_login
- method: PUT
- request params:
  - openid string
  - sid string
- request data:
  - avatar string
  - gender int
  - nickname string
- response status:
  - 200 success
  - 201 success new user
  - 401 auth check fails
  - 500 failure
```

2. 获取用户所属小组、创建者

```
# Get one group
- route: /group/:groupID
- method: GET
- request params:
  - openid string
  - sid string
- respnose data:
  - group Group
- response status:
  - 200 success
  - 401 auth check fails
  - 500 failure

# Get group by userID
- route: /group/user
- method: GET
- request params:
  - openid stirng
  - sid string
  - user_id string
- response data
  - group []Group
- response status:
  - 200 success
  - 401 auth check fails
  - 403 not allowed
  - 500 failure
```

3. 添加小组
```
# Create group
- route: /group
- method: PUT
- request params:
  - openid string
  - sid string
- request data:
  - name string
- response data:
  - group Group
  - err string # description on 500
- response status:
  - 201 success
  - 401 auth check fails
  - 403 not allowed
  - 500 failure

# Update group; get into the group
- route: /group/:groupID
- method: POST
- request params:
  - openid string
  - sid string
- request data:
  - name string
  - user_ids []string
- response data:
  - group Group
  - err string # description on 500
- response status:
  - 200 success
  - 201 success group changed
  - 401 auth check fails
  - 403 not allowed
  - 500 failure



# Delete group
- route: /group/:groupID
- method: DELETE
- request params:
  - openid string
  - sid string
- response status:
  - 200 success
  - 401 auth check fails
  - 403 not allowed
  - 500 failure
```


# 首页2（日历页）
1. 获取任务名、任务相关人员、任务类型、状态——显示未完成任务
```
# Get one task
- route: /task/:taskID/short
- method: GET
- request params:
  - openid string
  - sid string
- response data:
  - task:
    - name string
    - people []User
    - type int
    - status int
- response status:
  - 200 success
  - 401 auth check fails
  - 403 not allowed
  - 500 failure

```

2. 获取任务发布者、任务内容、DDL——显示具体任务
```
# Get one task
- route: /task/:taskID
- method: GET
- request params:
  - openid string
  - sid string
- response data:
  - task Task # computed field `done`
- response status:
  - 200 success
  - 401 auth check fails
  - 403 not allowed
  - 500 failure
```

# 小组页
1. 获取小组下所有任务名称、任务发布者、DDL——罗列小组内所有任务

```
# Get tasks by groupID
- route: /task/group
- method: GET
- request params:
  - openid string
  - sid string
  - group_id int
- response data:
  - task []Task # computed field `done`
- response status:
  - 200 success
  - 401 auth check fails
  - 403 not allowed
  - 500 failure
```

2. 获取任务相关人员、任务状态——标识未完成任务
3. 添加任务
```
# Create task
- route: /task
- method: PUT
- request params:
  - openid string
  - sid string
- request data:
  - group_id int
  - user_ids []string
  - name string
  - type int
  - leader_id # tbd
  - start_date Date: # easier to use then built-in date type; or timestamp I guess
    - year int
    - month int
    - day int  
  - end_date Date
  - description string
- response data:
    - task Task 
- response status:
  - 200 success
  - 201 success created
  - 400 wrong request format
  - 401 auth check fails
  - 403 not allowed
  - 500 failure
```

# 任务页
1. 获取任务名称、发布者、内容、DDL、相关人员、任务类型——显示具体任务及相关人员，确认Status修改权限

```
# Update a task; only update meta value; contents update would be implemented in next sprint
- route: /task/:taskID/meta
- method: POST
- request params:
  - openid string
  - sid string
- request data:
  - name string
  - start_date Date
  - end_date Date
  - readonly bool
  - description string
  - done bool
- response status:
  - 200 success
  - 201 success updated
  - 400 wrong request format
  - 401 auth check fails
  - 403 not allowed
  - 500 failure


# Delete task
- route: /tasks/:taskID
- method: DELETE
- request params:
  - openid string
  - sid string
- response status:
  - 200 success
  - 201 success deleted
  - 401 auth check fails
  - 403 not allowed
  - 500 failure
```
> mark
 
2. 修改任务属性（具体内容、相关人员、DDL、任务类型、任务状态）—获取修改记录

```
# record history is an advanced feature
// TODO: Design form of data for tasks modification records (in frontend, we want to display sth. like `git-diff` or `color-diff` for each modification)
```

3. （获取用户信息，添加评论）

```
# comments are advanced feature
- Get comments for one task
route: /tasks/:taskID/comments
method: GET
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




#
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

# 
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
