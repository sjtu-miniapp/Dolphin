package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sjtu-miniapp/dolphin/service/database/model"
	"log"
)

// REF: https://www.mindbowser.com/golang-go-with-gorm/
func DbConn(User, Password, Host, Db string, Port int) (*gorm.DB, error) {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		User, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// dont add s for tables
	db.SingularTable(true)
	DbSetup(db)
	return db, nil
}

func DbSetup(db *gorm.DB) {
	db = db.Debug()
	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
	}
	if !db.HasTable(&model.Group{}) {
		db.CreateTable(&model.Group{})
	}
	if !db.HasTable(&model.Task{}) {
		db.CreateTable(&model.Task{})
	}
	if !db.HasTable(&model.UserTask{}) {
		db.CreateTable(&model.UserTask{})
	}
	if !db.HasTable(&model.UserGroup{}) {
		db.CreateTable(&model.UserGroup{})
	}
	db.AutoMigrate(&model.User{}, &model.Task{},
		&model.Group{}, &model.UserTask{}, &model.UserGroup{})
	db.Model(&model.User{}).AddForeignKey("`self_group_id`",
		"`group`(id)", "SET NULL", "CASCADE")
	db.Model(&model.UserGroup{}).AddForeignKey("`user_id`",
		"`user`(id)", "CASCADE", "CASCADE")
	db.Model(&model.UserGroup{}).AddForeignKey("`group_id`",
		"`group`(id)", "CASCADE", "CASCADE")
	db.Model(&model.Group{}).AddForeignKey("`creator_id`",
		"`user`(id)", "CASCADE", "CASCADE")
	db.Model(&model.Task{}).AddForeignKey("`group_id`",
		"`group`(id)", "CASCADE", "CASCADE")
	db.Model(&model.Task{}).AddForeignKey("`leader_id`",
		"`user`(id)", "CASCADE", "CASCADE")
	db.Model(&model.Task{}).AddForeignKey("`publisher_id`",
		"`user`(id)", "CASCADE", "CASCADE")
	db.Model(&model.UserTask{}).AddForeignKey("`user_id`",
		"`user`(id)", "CASCADE", "CASCADE")
	db.Model(&model.UserTask{}).AddForeignKey("`task_id`",
		"`task`(id)", "CASCADE", "CASCADE")
}
