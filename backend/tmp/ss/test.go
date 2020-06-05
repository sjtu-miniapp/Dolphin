package main

import (
	"github.com/sjtu-miniapp/dolphin/service/database"
	"github.com/sjtu-miniapp/dolphin/service/database/model"
)

func main() {
	sqldb, _ := database.DbConn("root", "610878",
		"localhost", "test", 3306, 1)
	var count int32
	_ = sqldb.Model(&model.Task{}).Where(
		"group_id = ?", 30).Count(&count).Error
	println(count)
}
