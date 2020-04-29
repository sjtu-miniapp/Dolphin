package user

import (
	"context"
	"strings"
)

func (UserService) Hello(_ context.Context, name string) (string, error) {
	var message string
	if strings.EqualFold(name, "hello") {
		message = name
	} else {
		message = "world"
	}
	return message, nil
}
