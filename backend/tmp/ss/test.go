package main

import (
	"github.com/sjtu-miniapp/dolphin/service/database"
)
type authVal struct {
	Openid string `json:"openid"`
	Sid    string `json:"sid"`
}
func main() {
	redisdb, _ := database.InitRedis("121.199.33.44", "610878")
	database.Set(redisdb, 0, "test", authVal{Openid: "test", Sid: "test"})

}
