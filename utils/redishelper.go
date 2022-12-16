package utils

import (
	"fmt"
	config "rishabhjha12-hub/configRedis"
	"time"
)

// redis setter
func SET_FROM_REDIS(key string, value string, time time.Duration) {

	//reference = https://tutorialedge.net/golang/go-redis-tutorial/

	fmt.Println("key=", key, " value=", value)

	err := config.Client.Set(key, value, time).Err()

	if err != nil {
		fmt.Println(err)
	}

}

// redis setter array
func SET_FROM_REDIS_ARRAY(key string, value interface{}, time time.Duration) {

	//reference = https://tutorialedge.net/golang/go-redis-tutorial/

	fmt.Println("key=", key, " value=", value)

	err := config.Client.Set(key, value, time).Err()

	if err != nil {
		fmt.Println(err)
	}

}

// redis getter
func GET_FROM_REDIS(key string) string {

	//reference = https://tutorialedge.net/golang/go-redis-tutorial/
	keyVal, err := config.Client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}

	println(key, "----=", keyVal, "=-----")
	return keyVal
}
