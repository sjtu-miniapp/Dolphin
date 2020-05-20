package main


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


func DbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser,Password, Host, Port, Db )
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}
	// dont add s for tables
	db.SingularTable(true)
	return db
}

func main() {
	db := DbConn("root", "610878","localhost", "test", 3306)
	defer db.Close()
	db.AutoMigrate(&User{})
	db.Create(&User{
		Id:          "",
		Name:        "",
		Gender:      0,
		Avatar:      "",
		SelfGroupId: 0,
	})
	var user User
	// 为`name`列添加索引`idx_user_name`
	db.Model(&User{}).AddIndex("idx_user_name", "name")

	// 为`name`, `age`列添加索引`idx_user_name_age`
	db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")

	// 添加唯一索引
	db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")

	// 为多列添加唯一索引
	db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")

	// 删除索引
	db.Model(&User{}).RemoveIndex("idx_user_name")
	db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
	// id = 1
	db.First(&user, 1)
	// name = haha
	db.First(&user, "name = ?", "haha")
	// update name
	db.Model(&user).Update("name", "hahaha")

	db.Where()
}

