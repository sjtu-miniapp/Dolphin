package main

import (
	"fmt"
	"github.com/sjtu-miniapp/dolphin/service/database"
	"github.com/sjtu-miniapp/dolphin/service/database/model"
	"log"
)

func main() {
	db, err := database.DbConn("root", "610878", "localhost", "test", 3306, 1)
	if err != nil {
		log.Fatal(err)
	}
	//user := model.User{Id: "andy"}
	//err = db.Find(&user).Error
	//fmt.Println(err)
	//fmt.Println(user)
	//_ = rand.New(
	//	rand.NewSource(time.Now().UnixNano()))
	//name := "ahelsddaaadasdfalddffdadso" + strconv.Itoa(rand.Intn(1000))
	//// put user
	//user = model.User{Id: name, SelfGroupId: nil}
	//
	//db.Create(&user)
	//group := model.Group{
	//	Name:      new(string),
	//	CreatorId: name,
	//	Type:      0,
	//}
	//// put group
	//db.Create(&group)
	//db.Model(&user).Update("SelfGroupId", group.Id)
	////db.Find(&user).Update("SelfGroup", group)
	////fmt.Println(*user.SelfGroup, *user.SelfGroupId)
	//_ = db.Preload("SelfGroup").Find(&user).Error
	//db.Preload("Groups").Find(&user)
	//fmt.Println(user.Groups[0].Id)
	group := model.Group{
		Id:        6,
	}
	db.Preload("Users").Find(&group)
	fmt.Println(group.Users[0].Name)

}

// get
//	// id = 1
//	db.First(&user, 1)
//	// name = haha
//	db.First(&user, "name = ?", "haha")
//	db.Find(&user)

// Update with column names, not attribute names
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

//	// Using Find()
//	db.Find(&user).Update("Address", "San Diego")

//	// Batch Update
//	db.Table("user_models").Where("address = ?", "california").Update("name", "Walker")
//	// update if given a primary key, or insert, return primary key
//	db.Save()

//	db.Find(&user)
//	user.Address = "Brisbane"
//	db.Save(&user)


//func delete()  {

//	db.Table("user_models").Where("address= ?", "San Diego").Delete(&UserModel{})
//	//Find the record and delete it
//	db.Where("address=?", "Los Angeles").Delete(&UserModel{})
//
//	// Select all records from a model and delete all
//	db.Model(&UserModel{}).Delete(&UserModel{})
//}


//func get() {
//	// Get first record, order by primary key
//	db.First(&user)
//	// Get last record, order by primary key
//	db.Last(&user)
//	// Get all records
//	db.Find(&users)
//	// Get record with primary key (only works for integer primary key)
//	db.First(&user, 10)

//	db.Where("address = ?", "Los Angeles").First(&user)
//	//SELECT * FROM user_models WHERE address=’Los Angeles’ limit 1;
//	db.Where("address = ?", "Los Angeles").Find(&user)
//	//SELECT * FROM user_models WHERE address=’Los Angeles’;


//	db.Where("address <> ?", "New York").Find(&user)
//	//SELECT * FROM user_models WHERE address<>’Los Angeles’;

// IN
//	db.Where("name in (?)", []string{"John", "Martin"}).Find(&user)
//
//	// LIKE
//	db.Where("name LIKE ?", "%ti%").Find(&user)
//
//	// AND
//	db.Where("name = ? AND address >= ?", "Martin", "Los Angeles").Find(&user)
//}

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
