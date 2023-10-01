package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func CreateRedisClient() {

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	if _, err := client.Ping(context.TODO()).Result(); err != nil {
		log.Fatal(err)
	}
}

func GetRedisClient() *redis.Client {
	return client
}

func CloseRedisClient() error {
	return client.Close()
}
