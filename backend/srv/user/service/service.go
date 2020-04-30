package service

import (
	"context"
	"database/sql"
	"errors"
)

var (
	Err1 = errors.New("first error")
)

type Service interface {
	Hello(ctx context.Context, name string) (string, error)
}

type UserService struct {
	Db *sql.DB
}

