package main

import (
	"context"
	"fmt"
	auth "github.com/sjtu-miniapp/dolphin/service/auth/pb"
	"github.com/sjtu-miniapp/dolphin/service/group/pb"
	"strconv"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Group struct{}

func Router(base string) *gin.Engine {
	group := new(Group)
	router := gin.Default()
	g := router.Group(base)

	g.GET("/:group_id", group.GetGroup)
	// actually user_id
	g.GET("", group.GetGroupByUser)
	g.PUT("", group.CreateGroup)
	g.POST("/:group_id", group.UpdateGroup)
	g.DELETE("/:group_id", group.DeleteGroup)
	g.PUT("/:group_id/member", group.AddGroupMember)
	router.Use(cors.Default())
	return router
}

func inGroup(userId string, groupId int32) bool {
	resp, err := srv.UserInGroup(context.TODO(), &pb.UserInGroupRequest{
		UserId:  &userId,
		GroupId: &groupId,
	})
	if err != nil || !*resp.Ok {
		return false
	}
	return true
}

func checkAuth(c *gin.Context) error {
	openid := c.Query("openid")
	sid := c.Query("sid")
	res, err := authSrv.CheckAuth(context.TODO(),
		&auth.CheckAuthRequest{
			Openid: &openid,
			Sid:    &sid,
		},
	)
	if err != nil {
		return err
	} else if !*res.Ok {
		return fmt.Errorf("auth check fail")
	}
	return nil
}

/*
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
*/
func (g Group) GetGroup(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	openid := c.Query("openid")
	groupId, err := strconv.Atoi(c.Param("group_id"))
	gid := int32(groupId)
	if err != nil {
		c.JSON(400, err)
		return
	}
	if !inGroup(openid, gid) {
		c.JSON(403, "user is not in the group")
		return
	}
	resp, err := srv.GetGroup(context.TODO(), &pb.GetGroupRequest{
		Id: &gid,
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, resp)
}

/*
# Get group by userID
- route: /group/user
- method: GET
- request params:
    - openid stirng
    - sid string

- response data
  - group []Group
- response status:
    - 200 success
    - 401 auth check fails
    - 500 failure
*/
func (g Group) GetGroupByUser(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	openid := c.Query("openid")
	resp, err := srv.GetGroupByUserId(context.TODO(), &pb.GetGroupByUserIdRequest{
		UserId: &openid,
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, resp.Groups)
}

/*
# Create group
- route: /group
- method: PUT
- request params:
    - openid string
    - sid string
- request data:
    - name string
    - user_ids []string
- response data:
    - group Group
- response status:
    - 201 success
    - 401 auth check fails
    - 403 not allowed
    - 500 failure
*/
func (g Group) CreateGroup(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	openid := c.Query("openid")
	var data struct {
		Name    string   `json:"name"`
		UserIds []string `json:"user_ids"`
	}
	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(400, err)
		return
	}
	ty := int32(0)
	resp, err := srv.CreateGroup(context.TODO(), &pb.CreateGroupRequest{
		Name:      &data.Name,
		CreatorId: &openid,
		Type:      &ty,
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	_, err = srv.AddUser(context.TODO(), &pb.AddUserRequest{
		GroupId: resp.Id,
		UserIds: data.UserIds,
	})
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(201, resp.Id)
	}
}

/*
# Update group
- route: /group/:group_id
- method: POST
- request params:
    - openid string
    - sid string
- request data:
    - id int
    - name string
- response status:
    - 200 success
    - 201 success group changed
    - 401 auth check fails
    - 403 not allowed
    - 500 failure
*/
func (g Group) UpdateGroup(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	openid := c.Query("openid")
	var data struct {
		Id   int32 `json:"id"`
		Name string `json:"name"`
	}
	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(400, err)
		return
	}

	flag := inGroup(openid, data.Id)
	if !flag {
		c.JSON(403, "user not in group")
		return
	}

	_, err = srv.UpdateGroup(context.TODO(), &pb.UpdateGroupRequest{
		Id:   &data.Id,
		Name: &data.Name,
	})

	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, "updated")
}

/*
# Delete group
- route: /group/:groupID
- method: DELETE
- request params:
    - openid string
    - sid string
- request data:
    - id int
- response status:
    - 200 success
    - 401 auth check fails
    - 403 not allowed
    - 500 failure
*/
func (g Group) DeleteGroup(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	openid := c.Query("openid")
	var data struct {
		Id int32 `json:"id"`
	}
	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(400, err)
		return
	}
	r, err := srv.GetGroup(context.TODO(), &pb.GetGroupRequest{
		Id: &data.Id,
	})
	if err != nil {
		c.JSON(403, err)
		return
	}
	if *r.CreatorId != openid {
		c.JSON(403, "not allowed")
	}

	_, err = srv.DeleteGroup(context.TODO(), &pb.DeleteGroupRequest{
		Id: &data.Id,
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, "deleted")
}

func (g Group) AddGroupMember(c *gin.Context) {
	err := checkAuth(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	openid := c.Query("openid")
	id_, err := strconv.Atoi(c.Param("group_id"))
	if err != nil {
		c.JSON(500, err)
		return
	}
	id := int32(id_)
	var data struct {
		Action int32 `json:"add"`
		Members []string `json:"members"`
	}
	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(400, err)
		return
	}
	r, err := srv.GetGroup(context.TODO(), &pb.GetGroupRequest{
		Id:  &id,
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	if *r.CreatorId != openid {
		c.JSON(403, "not allowed")
		return
	}
	_, err = srv.AddGroupMembers(context.TODO(), &pb.AddGroupMembersRequest{
		Id:                   &id,
		Members:              data.Members,
		Action:               &data.Action,
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, "added")
}
