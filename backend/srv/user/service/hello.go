package service

import (
	"context"
	"github.com/sjtu-miniapp/dolphin/user/db"
)

func (us UserService) Hello(ctx context.Context, name string) (string, error) {
	err := db.UpdateName(ctx, us.Db, name, "hello")
	return "", err
}
