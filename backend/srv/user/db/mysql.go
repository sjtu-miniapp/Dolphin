package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Name   	string  `json:"name"`
	Message string `json:"message"`
}

func GetName(ctx context.Context, db *sql.DB, name string) (string, error) {
	sql1 := fmt.Sprintf("SELECT * FROM `test` WHERE `name` = '%s'", name)
	log.Println(sql1)
	rows := db.QueryRowContext(ctx, sql1)
	log.Println(rows)
	var user User
	err := rows.Scan(&user.Name, &user.Message)
	if err != nil {
		return "", err
	}
	//r, err := db.QueryContext(ctx, sql)
	//for r.Next()
	return user.Message, err

}

func DeleteName(ctx context.Context, db *sql.DB, name string) error {
	sql1 := fmt.Sprintf("DELETE FROM `test` WHERE `name` = '%s'", name)
	log.Println(sql1)
	_, err := db.ExecContext(ctx, sql1)
	return err
	// number of rows affected
	//result.RowsAffected()
	// multiple queries concurrently execute
	//db.Prepare()
}

func InsertName(ctx context.Context, db *sql.DB, name, message string) error {
	sql1 := fmt.Sprintf("INSERT INTO `test` VALUES('%s', '%s') ", name, message)
	log.Println(sql1)
	_, err := db.ExecContext(ctx, sql1)
	return err
}

func UpdateName(ctx context.Context, db *sql.DB, name, message string) error {
	sql1 := fmt.Sprintf("UPDATE `test` SET `name` = '%s',"+
		"`message` = '%s' WHERE `name` = '%s' or `message` = '%s' ", name, message, name, message)
	log.Println(sql1)
	_, err := db.ExecContext(ctx, sql1)
	return err
}
