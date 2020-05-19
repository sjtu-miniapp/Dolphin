package main

import (
	"context"
	"fmt"
	"github.com/sjtu-miniapp/dolphin/service/auth/pb"
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Auth struct{}

func Router(base string) *gin.Engine {
	auth := new(Auth)
	router := gin.Default()
	g := router.Group(base)
	g.POST("/on_login", auth.OnLogin)
	g.PUT("/after_login", auth.AfterLogin)

	router.Use(cors.Default())
	return router
}

/*
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
*/
// for uri: c.Param
// for params: c.Query, c.DefaultQuery
// for post form: c.PostForm
// for data: c.BindJSON
func (s *Auth) OnLogin(c *gin.Context) {
	//onlogin
	code := c.Query("code")
	//c.GetRawData()
	if code == "" {
		c.JSON(400, fmt.Errorf("no code for query"))
		return
	}
	response, err := srv.OnLogin(context.TODO(), &pb.OnLoginRequest{
		Code: code,
	})
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, response)
	}
}

/*
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
*/

func (s *Auth) AfterLogin(c *gin.Context) {
	openid := c.Query("openid")
	sid := c.Query("sid")
	var data struct {
		Avatar   string `json:"avatar"`
		Gender   int    `json:"gender"`
		Nickname string `json:"nickname"`
	}
	err := c.BindJSON(&data)
	if openid == "" || sid == "" || err != nil {
		c.JSON(400, err)
		return
	}
	res, err := srv.CheckAuth(context.TODO(), &pb.CheckAuthRequest{
		Openid: openid,
		Sid:    sid,
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	if res.Ok != true {
		c.JSON(401, "check auth failed")
		return
	}

	resp, err := srv.PutUser(context.TODO(), &pb.PutUserRequest{
		Openid: openid,
		Name:   data.Nickname,
		Gender: int32(data.Gender),
		Avatar: data.Avatar,
	})
	if err != nil {
		c.JSON(500, err)
	} else if resp.Err == 1 {
		c.JSON(200, resp)
	} else {
		c.JSON(201, resp)
	}
}
