package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)
type Date struct {
	Year int `json:"year"`
	Month int `json:"month"`
	Day int `json:"day"`
}
/*
default:
	string:""
	int:0
	array: null
	struct: all default(but could be nil)

*/
func handle(c *gin.Context) {
	var data struct {
		GroupId *int `json:"group_id"`
		UserIds []*string `json:"user_ids"`
		Name *string `json:"name, omitempty"`
		Type int 	`json:"type"`
		LeaderId string `json:"leader_id, omitempty"`
		StartDate Date `json:"start_date, omitempty"`
		EndDate Date `json:"end_date, omitempty"`
		Description string `json:"description, omitempty"`
	}

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, data)
	fmt.Println(data)
}

func main() {
	service := web.NewService()
	service.Init()
	router := gin.Default()
	router.POST("/", handle)
	service.Handle("/", router)
	service.Run()
}