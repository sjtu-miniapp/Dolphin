package main

import (
	"github.com/sjtu-miniapp/dolphin/service/database"
	"github.com/sjtu-miniapp/dolphin/service/database/model"
)

func main() {
	sqldb, _ := database.DbConn("root", "610878",
		"localhost", "test", 3306, 1)
	var count int32

	_ = sqldb.Model(&model.Group{
		Id: 1,
	}).Update("updated_at", 0).Error
	println(count)
}
