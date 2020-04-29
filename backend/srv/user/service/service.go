package service

import (
	"context"
	"errors"
)

var (
	Err1 = errors.New("first error")
)

type Service interface {
	Hello(ctx context.Context, name string) (string, error)
}

type UserService struct {}
