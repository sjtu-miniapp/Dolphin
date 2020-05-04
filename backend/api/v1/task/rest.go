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

	g.GET("/task", task.OnLogin)

	router.Use(cors.Default())
	return router
}

func (s *Task) OnLogin(c *gin.Context) {
	//// parse request
	//id, err := strconv.ParseInt(c.Query("id"), 0, 64)
	//// get response
	//response, err := srv.GetTask(context.TODO(), &pb.GetTaskRequest{
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
