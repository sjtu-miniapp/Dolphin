package database


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sjtu-miniapp/dolphin/service/database/model"
	"log"
)

// REF: https://www.mindbowser.com/golang-go-with-gorm/
func DbConn(MyUser, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser,Password, Host, Port, Db )
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
	}
	// dont add s for tables
	db.SingularTable(true)
	DbSetup(db)
	return db
}

func DbSetup(db *gorm.DB) {
	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
	}
	if !db.HasTable(&model.Group{}) {
		db.CreateTable(&model.Group{})
	}
	if !db.HasTable(&model.Task{}) {
		db.CreateTable(&model.Task{})
	}
	db.Debug().AutoMigrate(&model.User{}, &model.Task{}, model.Group{})
	db.Model(&model.User{}).AddForeignKey("self_group_id",
		"group(id)", "SET NULL", "CASCADE")
	db.Model(&model.UserGroup{}).AddForeignKey("user_id",
		"user(id)", "CASCADE", "CASCADE")
	db.Model(&model.UserGroup{}).AddForeignKey("group_id",
		"group(id)", "CASCADE", "CASCADE")
	db.Model(&model.Group{}).AddForeignKey("creator_id",
		"user(id)", "CASCADE", "CASCADE")
	db.Model(&model.Task{}).AddForeignKey("group_id",
		"group(id)", "CASCADE", "CASCADE")
	db.Model(&model.Task{}).AddForeignKey("leader_id",
		"user(id)", "CASCADE", "CASCADE")
	db.Model(&model.Task{}).AddForeignKey("publisher_id",
		"user(id)", "CASCADE", "CASCADE")
	db.Model(&model.UserTask{}).AddForeignKey("user_id",
		"user(id)", "CASCADE", "CASCADE")
	db.Model(&model.UserTask{}).AddForeignKey("task_id",
		"task(id)", "CASCADE", "CASCADE")

}


//func main() {
//
//	db := DbConn("root", "610878","localhost", "test", 3306)
//	defer db.Close()
//
//	db.AutoMigrate(Event{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
//	db.AutoMigrate(&User{})
//	db.Create(&User{
//		Id:          "",
//		Name:        "",
//		Gender:      0,
//		Avatar:      "",
//		SelfGroupId: 0,
//	})
//	var user User
//	// 为`name`列添加索引`idx_user_name`
//	db.Model(&User{}).AddIndex("idx_user_name", "name")
//
//	// 为`name`, `age`列添加索引`idx_user_name_age`
//	db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")
//
//	// 添加唯一索引
//	db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")
//
//	// 为多列添加唯一索引
//	db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")
//
//	// 删除索引
//	db.Model(&User{}).RemoveIndex("idx_user_name")
//	db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
//	// id = 1
//	db.First(&user, 1)
//	// name = haha
//	db.First(&user, "name = ?", "haha")
//	db.Find(&user)
//	// Update with column names, not attribute names
//	db.Model(&user).Update("Name", "hahaha")
//
//	db.Model(&user).Updates(
//		map[string]interface{}{
//			"Name": "Amy",
//			"Address": "Boston",
//		})
//	// UpdateColumn()
//	db.Model(&user).UpdateColumn("Address", "Phoenix")
//	db.Model(&user).UpdateColumns(
//		map[string]interface{}{
//			"Name": "Taylor",
//			"Address": "Houston",
//		})
//
//	db.Where()
//	db.Create()
//	// Using Find()
//	db.Find(&user).Update("Address", "San Diego")
//
//	// Batch Update
//	db.Table("user_models").Where("address = ?", "california").Update("name", "Walker")
//	// update if given a primary key, or insert, return primary key
//	db.Save()
//
//}
//
//func createInsert() {
//user:=&UserModel{Name:"John",Address:"New York"}
//	db.Create(user)
//
//	Internally it will create the query like
//	INSERT INTO `user_models` (`name`,`address`) VALUES ('John','New York')
//
//	//You can insert multiple records too
//	var users []UserModel = []UserModel{
//		UserModel{name: "Ricky",Address:"Sydney"},
//		UserModel{name: "Adam",Address:"Brisbane"},
//		UserModel{name: "Justin",Address:"California"},
//	}
//
//	for _, user := range users {
//		db.Create(&user)
//	}
//}
//func update() {
//	user:=&UserModel{Name:"John",Address:"New York"}
//	// Select, edit, and save
//	db.Find(&user)
//	user.Address = "Brisbane"
//	db.Save(&user)
//
//	// Update with column names, not attribute names
//	db.Model(&user).Update("Name", "Jack")
//
//	db.Model(&user).Updates(
//		map[string]interface{}{
//			"Name": "Amy",
//			"Address": "Boston",
//		})
//
//	// UpdateColumn()
//	db.Model(&user).UpdateColumn("Address", "Phoenix")
//	db.Model(&user).UpdateColumns(
//		map[string]interface{}{
//			"Name": "Taylor",
//			"Address": "Houston",
//		})
//	// Using Find()
//	db.Find(&user).Update("Address", "San Diego")
//
//	// Batch Update
//	db.Table("user_models").Where("address = ?", "california").Update("name", "Walker")
//}
//
//func delete()  {
//	// Select records and delete it
//	db.Table("user_models").Where("address= ?", "San Diego").Delete(&UserModel{})
//
//	//Find the record and delete it
//	db.Where("address=?", "Los Angeles").Delete(&UserModel{})
//
//	// Select all records from a model and delete all
//	db.Model(&UserModel{}).Delete(&UserModel{})
//}
//
//func get() {
//	// Get first record, order by primary key
//	db.First(&user)
//	// Get last record, order by primary key
//	db.Last(&user)
//	// Get all records
//	db.Find(&users)
//	// Get record with primary key (only works for integer primary key)
//	db.First(&user, 10)
//
//	Query with Where() [some SQL functions]
//
//	db.Where("address = ?", "Los Angeles").First(&user)
//	//SELECT * FROM user_models WHERE address=’Los Angeles’ limit 1;
//
//	db.Where("address = ?", "Los Angeles").Find(&user)
//	//SELECT * FROM user_models WHERE address=’Los Angeles’;
//
//	db.Where("address <> ?", "New York").Find(&user)
//	//SELECT * FROM user_models WHERE address<>’Los Angeles’;
//
//	// IN
//	db.Where("name in (?)", []string{"John", "Martin"}).Find(&user)
//
//	// LIKE
//	db.Where("name LIKE ?", "%ti%").Find(&user)
//
//	// AND
//	db.Where("name = ? AND address >= ?", "Martin", "Los Angeles").Find(&user)
//}
//
//func tx()  {
//	tx := db.Begin()
//	err := tx.Create(&user).Error
//	if err != nil {
//		tx.Rollback()
//	}
//	tx.Commit()
//}
//
//func as()  {
//	db.Model(&places).Association("town").Find(&places.Town)
}