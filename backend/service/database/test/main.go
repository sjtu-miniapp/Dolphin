package main

import "github.com/sjtu-miniapp/dolphin/service/database"

func main() {
	db := database.DbConn("root", "610878", "localhost", "test", 3306)
	database.DbSetup(db)

}
