package main

import (
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Task struct{}



func Router(base string) *gin.Engine {
	task := new(Task)
	router := gin.Default()
	g := router.Group(base)

	g.GET("/task/:task_id/meta", task.GetTaskMeta)
	g.GET("/task/:task_id/workers", task.GetTaskWorker)
	g.GET("/task/:task_id", task.GetTask)
	g.GET("/task/group", task.GetTaskByGroup)
	g.PUT("/task", task.CreateTask)
	g.POST("/task/:task_id/meta", task.UpdateTaskMeta)
	g.DELETE("/task/:task_id", task.DeleteTask)
	g.POST("/task/:task_id", task.UpdateTask)
	router.Use(cors.Default())
	return router
}
/*
# Get one task
- route: /task/:taskId/meta
- method: GET
- request params:
    - openid string
    - sid string
- response data:
    - name string
    - type int
    - done bool
    - group_id int
    - publisher_id string
    - leader_id string
    - start_date Date
    - end_date Date
    - readonly bool
    - type int
    - description string
- response status:
    - 200 success
    - 401 auth check fails
    - 403 not allowed
    - 500 failure
 */
func (t Task) GetTaskMeta(context *gin.Context) {
	srv.GetTaskMeta()
}
/*
# Get worker of the task
- route: /task/:taskId/workers
- method: GET
- request params:
    - openid string
    - sid string
- response data:
    - workers []User
- response status:
    - 200 success
    - 401 auth check fails
    - 403 not allowed
    - 500 failure
 */
func (t Task) GetTaskWorker(context *gin.Context) {
	srv.GetTaskPeolple()
}


/*
# Get task content # tbd # advanced
- route: /task/:taskID
- method: GET
  - request params:
      - openid string
      - sid string
  - response data:

  - response status:
      - 200 success
      - 401 auth check fails
      - 403 not allowed
      - 500 failure
 */
func (t Task) GetTask(context *gin.Context) {
	panic("implement me")
}
/*
# Get tasks by groupID
- route: /task/group
- method: GET
- request params:
    - openid string
    - sid string
    - group_id int
- response data:
    - task []Task # meta
- response status:
    - 200 success
    - 401 auth check fails
    - 403 not allowed
    - 500 failure
 */
func (t Task) GetTaskByGroup(context *gin.Context) {
	srv.GetTaskMetaByGroupId()
}
/*
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
 */
func (t Task) CreateTask(context *gin.Context) {
	srv.CreateTask()
}
/*
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
 */
func (t Task) UpdateTaskMeta(context *gin.Context) {

}
/*
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

 */
func (t Task) DeleteTask(context *gin.Context) {

}

/*
# Update data content; advanced feature
- route: /task/:taskId
- method: POST
- request params:
    - openid string
    - sid string
...
 */
func (t Task) UpdateTask(context *gin.Context) {
	panic("implement me")
}

