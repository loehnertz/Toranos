package commons

import (
	"github.com/go-redis/redis"
	"github.com/micro/go-log"
)

const RedisHostAddress = "localhost:6379"
const RedisDatabaseId = 8

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
