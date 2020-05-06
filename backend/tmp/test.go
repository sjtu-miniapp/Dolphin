package main

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
//func handle(c *gin.Context) {
//	var data struct {
//		GroupId *int `json:"group_id"`
//		UserIds []*string `json:"user_ids"`
//		Name *string `json:"name, omitempty"`
//		Type int 	`json:"type"`
//		LeaderId string `json:"leader_id, omitempty"`
//		StartDate Date `json:"start_date, omitempty"`
//		EndDate Date `json:"end_date, omitempty"`
//		Description string `json:"description, omitempty"`
//	}
//
//	err := c.BindJSON(&data)
//	if err != nil {
//		c.JSON(500, err)
//		return
//	}
//	c.JSON(200, data)
//	fmt.Println(data)
//}
//func test1() {
//	service := web.NewService()
//	service.Init()
//	router := gin.Default()
//	router.POST("/", handle)
//	service.Handle("/", router)
//	service.Run()
//}

//func test2() {
//	db, _ := database.InitDb("root", "610878", "localhost", "dolphin")
//	rows, err := db.QueryContext(context.Background(),
//		"SELECT `id`, `creator_id`, `name` FROM `user_group` JOIN `group` ON `group_id`=`id` WHERE `user_id`= ?", "anofa")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		var data struct{
//			Id int `json:"id"`
//			CreatorId string `json:"creator_id"`
//			Name sql.NullString `json:"name"`
//		}
//		if rows.Next() {
//			err = rows.Scan(&data.Id, &data.CreatorId, &data.Name)
//		}
//		fmt.Println(err, data)
//	}
//}
func main() {
	var a []*int
	b := 1
	c := append(a, &b)
	print(len(c))
}
