package main

import (
	"context"
	"strconv"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/sjtu-miniapp/dolphin/service/group/pb"
)

type Group struct{}

func Router(base string) *gin.Engine {
	group := new(Group)
	router := gin.Default()
	g := router.Group(base)

	g.GET("/group", group.GetGroup)

	router.Use(cors.Default())
	return router
}

func (s *Group) GetGroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 0, 64)

	response, err := srv.GetGroup(context.TODO(), &pb.GetGroupRequest{
		Id: id,
	})

	// error response
	if err != nil {
		c.JSON(500, err)
	} else {
		// correct response
		c.JSON(200, response)
	}
}
