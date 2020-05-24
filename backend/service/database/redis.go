package database

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
	log.Printf("[redis]: get key %s\n", key)
	if err != nil {
		log.Println("[redis]: can't get the value")
		return err
	}
	err = json.Json2struct(val, s)
	return err
}


