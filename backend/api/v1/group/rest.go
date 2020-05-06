package main

import (
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Group struct{}

func Router(base string) *gin.Engine {
	group := new(Group)
	router := gin.Default()
	g := router.Group(base)

	g.GET("/group/{group_id}", group.GetGroup)
	g.GET("/group/user", group.GetGroupByUser)
	g.PUT("/group", group.CreateGroup)
	g.POST("/group/{group_id}", group.UpdateGroup)
	g.DELETE("/group/{group_id}", group.DeleteGroup)
	router.Use(cors.Default())
	return router
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
func (g Group) GetGroup(context *gin.Context) {
	srv.GetGroup()
	panic("implement me")
}
/*
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
 */
func (g Group) GetGroupByUser(context *gin.Context) {
	srv.GetGroup()
	panic("implement me")

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
- response data:
    - group Group
    - err string # description on 500
- response status:
    - 201 success
    - 401 auth check fails
    - 403 not allowed
    - 500 failure
 */
func (g Group) CreateGroup(context *gin.Context) {
	srv.create
	srv.adduser
	panic("implement me")

}
/*
# Update group
- route: /group/:groupID
- method: POST
- request params:
    - openid string
    - sid string
- request data:
    - id int
    - name string
- response data:
    - group Group
    - err string # description on 500
- response status:
    - 200 success
    - 201 success group changed
    - 401 auth check fails
    - 403 not allowed
    - 500 failure
 */
func (g Group) UpdateGroup(context *gin.Context) {
	panic("implement me")

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
    - name string
- response status:
    - 200 success
    - 401 auth check fails
    - 403 not allowed
    - 500 failure
 */
func (g Group) DeleteGroup(context *gin.Context) {
	panic("implement me")

}