package config

import (
	"rishabhjha12-hub/constants"

	"github.com/go-redis/redis"
)

// exporting redis client
var Client *redis.Client

// configure redis
// reference = https://tutorialedge.net/golang/go-redis-tutorial/
func RedisConfig() {
	client := redis.NewClient(&redis.Options{
		Addr:     constants.Redisadress, // host:port of the redis server
		Password: constants.RedisPass,   // no password set
		DB:       constants.Redis_db,    // use default DB
	})
	Client = client

}
