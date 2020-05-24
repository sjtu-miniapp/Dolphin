package main

import (
	"github.com/go-redis/redis"
	"github.com/sjtu-miniapp/dolphin/utils/json"
	"log"
	"time"
)

func InitRedis(host string, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: password,
	})
	err := client.Ping().Err()
	if err != nil {
		log.Println(err)
		log.Fatal("fail to connect to redis")
		return nil, err
	}
	return client, nil
}

func Set(c *redis.Client, ttl time.Duration, key string, value interface{}) error {
	jsonData := json.Struct2json(value)
	err := c.Set(key, jsonData, ttl).Err()
	if err != nil {
		log.Println("[redis]: set value failed", err.Error())
	} else {
		log.Printf("[redis]: set %s to %s\n", key, value)
	}
	return err
}

func Get(c *redis.Client, key string, s interface{}) error {
	val, err := c.Get(key).Result()
	if err != nil {
		log.Println("[redis]: can't get the value")
		return err
	}
	err = json.Json2struct(val, s)
	return err
}

func main() {
	redis, _ := InitRedis("121.199.33.44", "610878")
	var s string
	Get(redis, "87198d9900b2ab7cd8fdc177a94bc9189963a72fdd928a9663a60f8475e064c4", &s)
	print(s)
}
