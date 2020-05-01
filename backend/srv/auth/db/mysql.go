package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sjtu-miniapp/dolphin/auth/logger"
	"log"
)

type User struct {
	Id   string
	Name string
	Gender int
	SelfGroupId int
}

func GetUser(ctx context.Context, db *sql.DB, id string) (User, error) {
	sql1 := fmt.Sprintf("SELECT * FROM `test` WHERE `id = '%s'", id)
	logger.Log.Debug(sql1)
	rows, err := db.QueryContext(ctx, sql1)
	if err != nil {
		return User{}, err
	}
	var user User
	if rows.Next() {
		_ = rows.Scan(&user.Id, &user.Name, &user.Gender, &user.SelfGroupId)
	} else {
		return User{}, fmt.Errorf("no user found")
	}
	return user, nil
}


func InsertUser(ctx context.Context, db *sql.DB, id, name string, gender uint32) error {
	_, err :=  GetUser(ctx, db, id)
	if err != nil {
		return nil
	} else {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		{
			sql1 := fmt.Sprintf("INSERT INTO `user`(`id`, `name`, `gender`)"+
				" VALUES('%s', '%s', %d) ", id, name, gender)
			logger.Log.Debug(sql1)
			stmt, err := tx.PrepareContext(ctx, sql1)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
			_, err = stmt.ExecContext(ctx)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
		}
		{
			sql1 := fmt.Sprintf("INSERT INTO `group`(`creator_id`, `type`)"+
				" VALUES('%s', 'INDIVIDUAL') ", id)
			logger.Log.Debug(sql1)
			stmt, err := tx.PrepareContext(ctx, sql1)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
			result, err := stmt.ExecContext(ctx)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
			stmt, err = tx.PrepareContext(ctx, "SET foreign_key_checks = 0")
			_, _ = stmt.ExecContext(ctx)
			selfGroupId, _ := result.LastInsertId()
			sql2 := fmt.Sprintf("UPDATE `user` SET `self_group_id` = %d "+
				" WHERE `id` = '%s'", selfGroupId, id)
			stmt, err = tx.PrepareContext(ctx, sql2)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
			result, err = stmt.ExecContext(ctx)
			if err != nil {
				_ = tx.Rollback()
				return err
			}
			stmt, err = tx.PrepareContext(ctx, "SET foreign_key_checks = 1")
			_, _ = stmt.ExecContext(ctx)
		}

		err = tx.Commit()
		return err
	}


}

func UpdateName(ctx context.Context, db *sql.DB, name, message string) error {
	sql1 := fmt.Sprintf("UPDATE `test` SET `name` = '%s',"+
		"`message` = '%s' WHERE `name` = '%s' or `message` = '%s' ", name, message, name, message)
	log.Println(sql1)
	_, err := db.ExecContext(ctx, sql1)
	return err
}

