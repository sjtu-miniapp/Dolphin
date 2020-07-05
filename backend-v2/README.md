# TODO

basePrefix: /api/v2
baseQuery: openid={openid}&sid={sid}

## Home Page (Group in General)

prefix: {basePrefix}/group

### Get Group
url: {prefix}?{baseQuery}
method: GET
response data:
  {
    id: string;
    title: string;
    taskNumber: number;
    updateTime: number; // timestamp
  }[]

### Add Group
url: {prefix}?{baseQuery}
method: POST
request data: { groupName: string }
response status: 200, 409, 400, 403, ...

### Join Group
url: {prefix}/sharecode/:{sharecode}?{baseQuery}
method: POST
response data:
{
  groupID:string; groupName: string
}

## Group Detail Page

prefix: {basePrefix}/group

### Get Task Short List
url: {prefix}/:groupID/tasks?{baseQuery}
method: GET
response data:
{
  id: string;
  title: string;
  status: string;
  ddl: number; // timestamp
}

### Click Sharecode
url: {prefix}/sharecode/:groupID?{baseQuery}
method: GET
response data:
{
  sharecode: number;
  groupID: id;
}

### Create Task
url: {prefix}/:groupID/task?{baseQuery}
method: post
request data:
{
  title: string;
  type: string;
  ddl: number;
  description: string;
}
response status: 200, 409, 400, 403, ...

## Task Page

prefix: {basePrefix}/task

### Get Task
url: {prefix}/:taskID?{baseQuery}
method: GET
response data:
{
  title: string;
  publisher: string;
  createTime: string;
  type: string;
  ddl: string;
  status: string;
  members:
  {
    id: string;
    name: string;
    status: string;
  }[];
}

### Update Memeber Status
url: {prefix}/:taskID/member/:memberid?{baseQuery}
method: POST
request data: { status: string }
response: 200

### Update Task
url: {prefix}/:taskID?{baseQuery}
method: post
request data:
{
  title?: string;
  ddl?: string;
  desciption?: string;
  status?: string;
}

