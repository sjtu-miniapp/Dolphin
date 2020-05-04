package main

import (
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Auth struct{}

func Router(base string) *gin.Engine {
	auth := new(Auth)
	router := gin.Default()
	g := router.Group(base)

	g.GET("/auth", auth.OnLogin)

	router.Use(cors.Default())
	return router
}

func (s *Auth) OnLogin(c *gin.Context) {
	//// parse request
	//id, err := strconv.ParseInt(c.Query("id"), 0, 64)
	//// get response
	//response, err := srv.GetAuth(context.TODO(), &pb.GetAuthRequest{
	//	Id: id,
	//})
	//
	//// handle response
	//if err != nil {
	//	c.JSON(500, err)
	//} else {
	//	c.JSON(200, response)
	//}
}
