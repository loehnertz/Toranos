package common

import (
	"github.com/go-redis/redis"
	"github.com/micro/go-log"
)

func InitRedisClient(host string, password string, database int) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       database,
	})

	_, connectionError := client.Ping().Result()
	if connectionError != nil {
		log.Log(connectionError)
	}

	return
}
