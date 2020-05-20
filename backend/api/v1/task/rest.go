package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/util/log"
	pb2 "github.com/sjtu-miniapp/dolphin/service/group/pb"
	"github.com/sjtu-miniapp/dolphin/service/task/pb"
	"strconv"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	auth "github.com/sjtu-miniapp/dolphin/service/auth/pb"
)

type Task struct{}

func Router(base string) *gin.Engine {
	task := new(Task)
	router := gin.Default()
	g := router.Group(base)

	g.GET("/meta/:task_id", task.GetTaskMeta)
	g.GET("/workers/:task_id", task.GetTaskWorker)
	g.GET("/content/:task_id", task.GetTask)
	g.GET("/group", task.GetTaskByGroup)
	g.PUT("/", task.CreateTask)
	g.POST("/meta/:task_id", task.UpdateTaskMeta)
	g.DELETE("/:task_id", task.DeleteTask)
	g.POST("/:task_id", task.UpdateTask)
	router.Use(cors.Default())
	return router
}

func checkAuth(c *gin.Context) error {
	openid := c.Query("openid")
	sid := c.Query("sid")
	res, err := authSrv.CheckAuth(context.TODO(),
		&auth.CheckAuthRequest{
			Openid: openid,
			Sid:    sid,
		},
	)
	if err != nil {
		return err
	}
	if !res.Ok {
		return fmt.Errorf("auth check fail")
	}
	return nil
}

func inTask(userId string, taskId uint32) bool {

	resp, err := srv.GetTaskMeta(context.TODO(), &pb.GetTaskMetaRequest{
		Id:                   taskId,
	})
	if err != nil || resp.Meta == nil {
		return false
	}
	if resp.Meta.PublisherId == userId {
		return true
	}

	resp2, err := srv.UserInTask(context.TODO(), &pb.UserInTaskRequest{
		UserId: userId,
		TaskId: taskId,
	})
	if err != nil {
		log.Error(err)
		return false
	}
	return resp2.Ok
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
func (t Task) GetTaskMeta(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}

	openid := c.Query("openid")
	id, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	if !inTask(openid, uint32(id)) {
		c.JSON(403, "not allowed")
	}

	resp, err := srv.GetTaskMeta(context.TODO(), &pb.GetTaskMetaRequest{
		Id: uint32(id),
	})
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, resp.Meta)
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
func (t Task) GetTaskWorker(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	openid := c.Query("openid")
	id, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	if !inTask(openid, uint32(id)) {
		c.JSON(403, "not allowed")
	}
	resp, err := srv.GetTaskPeolple(context.TODO(), &pb.GetTaskPeopleRequset{
		Id: uint32(id),
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, resp)
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
func (t Task) GetTask(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, "auth check fail")
		return
	}
	c.JSON(555, "NOT IMPLEMENTED YET!!!!")
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
func (t Task) GetTaskByGroup(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	openid := c.Query("openid")
	groupId, err := strconv.Atoi(c.Query("group_id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	resp, err := groupSrv.UserInGroup(context.TODO(), &pb2.UserInGroupRequest{
		UserId:  openid,
		GroupId: uint32(groupId),
	})
	if err != nil {
		c.JSON(403, err)
		return
	}

	if !resp.Ok {
		c.JSON(403, "not allowed")
		return
	}
	resp2, err := srv.GetTaskMetaByGroupId(context.TODO(), &pb.GetTaskMetaByGroupIdRequest{
		GroupId: uint32(groupId),
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, resp2.Metas)
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

func (t Task) CreateTask(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, "auth check fail")
		return
	}
	var data struct {
		GroupId     int      `json:"group_id"`
		UserIds     []string `json:"user_ids"`
		Name        string   `json:"name"`
		Type        int      `json:"type"`
		LeaderId    string   `json:"leader_id, omitempty"`
		StartDate   string `json:"start_date, omitempty"`
		EndDate     string `json:"end_date, omitempty"`
		Description string   `json:"description, omitempty"`
		Readonly    bool     `json:"readonly"`
	}
	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(400, err)
		return
	}
	for _, v := range data.UserIds {
		rsp, err := groupSrv.UserInGroup(context.Background(), &pb2.UserInGroupRequest{
			UserId:               v,
			GroupId:              uint32(data.GroupId),
		})
		if err != nil {
			c.JSON(500, err)
			return
		}
		if !rsp.Ok {
			c.JSON(400, "user not in group")
			return
		}
	}

	openid := c.Query("openid")
	resp, err := srv.CreateTask(context.TODO(), &pb.CreateTaskRequest{
		GroupId:     int32(data.GroupId),
		Readonly:    data.Readonly,
		UserIds:     data.UserIds,
		Name:        data.Name,
		Type:        int32(data.Type),
		LeaderId:    data.LeaderId,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
		Description: data.Description,
		PublisherId: openid,
	})

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(201, resp.Id)
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
func (t Task) UpdateTaskMeta(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	openid := c.Query("openid")
	id, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		c.JSON(400, err)
		return
	}

	if !inTask(openid, uint32(id)) {
		c.JSON(403, "not allowed")
		return
	}

	resp, err := srv.GetTaskMeta(context.TODO(), &pb.GetTaskMetaRequest{
		Id: uint32(id),
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	if resp.Meta == nil {
		c.JSON(400, "no task found")
		return
	}

	if resp.Meta.Readonly && resp.Meta.PublisherId != openid {
		c.JSON(403, "not allowed")
		return
	}

	var data struct {
		GroupId     int      `json:"group_id"`
		Name        *string   `json:"name"`
		StartDate   *string `json:"start_date, omitempty"`
		EndDate     *string `json:"end_date, omitempty"`
		Description *string   `json:"description, omitempty"`
		Readonly 	*bool `json:"readonly"`
	}

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(400, err)
		return
	}
	if data.StartDate == nil {
		*data.StartDate = resp.Meta.StartDate
	}
	if data.EndDate == nil {
		*data.EndDate = resp.Meta.EndDate
	}
	if data.Name == nil {
		*data.Name = resp.Meta.Name
	}
	if data.StartDate == nil {
		*data.StartDate = resp.Meta.StartDate
	}
	if data.EndDate == nil {
		*data.EndDate = resp.Meta.EndDate
	}
	if data.Description == nil {
		*data.Description = resp.Meta.Description
	}
	if data.Readonly == nil {
		*data.Readonly = resp.Meta.Readonly
	}
	_, err = srv.UpdateTaskMeta(context.TODO(), &pb.UpdateTaskMetaRequest{
		Id:          uint32(id),
		Name:        *data.Name,
		StartDate:   *data.StartDate,
		EndDate:     *data.EndDate,
		Readonly:    *data.Readonly,
		Description: *data.Description,
	})
	if err != nil {
		c.JSON(500 ,err)
		return
	}
	c.JSON(201, "updated")
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
func (t Task) DeleteTask(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	openid := c.Query("openid")

	id, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		c.JSON(400, err)
		return
	}
	// FIXME: MAYBE DON'T ALLOW MEMBERS TO DELETE TASK
	if !inTask(openid, uint32(id)) {
		c.JSON(403, "not allowed")
		return
	}

	_, err = srv.DeleteTask(context.TODO(), &pb.DeleteTaskRequest{
		Id:                   uint32(id),
	})

	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, "deleted")
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
func (t Task) UpdateTask(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, "auth check fail")
		return
	}
	//openid := c.Query("openid")
	c.JSON(555, "NOT IMPLMENTED YET!")
}
