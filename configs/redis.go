package configs

import "github.com/go-redis/redis"

var Client *redis.Client

func GetRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
