package main

import (
	//"encoding/json"
	"fmt"
	configmqtt "rishabhjha12-hub/configMqtt"
	config "rishabhjha12-hub/configRedis"
	configsentry "rishabhjha12-hub/configSentry"
	"rishabhjha12-hub/constants"
	"rishabhjha12-hub/helper"
	"rishabhjha12-hub/utils"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	config.RedisConfig()
}
func TestGetterSetter(t *testing.T) {
	utils.SET_FROM_REDIS("keyss", "50", 100*time.Second)
	utils.GET_FROM_REDIS("keyss")
	assert.NotNil(t, "keyss")
	if assert.NotNil(t, "keys") {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "50", utils.GET_FROM_REDIS("keyss"))
	}

}
func TestConfigSentry(t *testing.T) {
	configsentry.SentryConfig()
}

func TestUnmarshalAndPrint(t *testing.T) {
	t.Run("testing unmarshalAndPrint", func(t *testing.T) {

		myjson := ` {
			"sender_mac_add":       "94:b9:7e:6b:76:1b",
			"receiver_mac": "11:11:11:11:11:11",
			"message_type_code":    8000,
			"message_body": {
					"epc_id":       "34161FA8203288AC05FF4500",
					"plaza_id":     "15"
			}
	}`

		val := helper.ParseHelper([]byte(myjson))
		assert.Equal(t, val, constants.Parkzap_Msg_Code, "jk")

		fmt.Println("*******======", val)
	})
}
func TestMqtt(t *testing.T) {
	configmqtt.MqttConfig()
}

// reference  https://github.com/alicebob/miniredis
func TestRediss(t *testing.T) {
	s := miniredis.RunT(t)

	// Optionally set some keys your code expects:
	s.Set("foo", "bar")
	s.HSet("some", "other", "key")

	// Run your code and see if it behaves.
	// An example using the redigo library from "github.com/gomodule/redigo/redis":
	c, err := redis.Dial("tcp", s.Addr())

	_, err = c.Do("SET", "foo", "bar")
	if err != nil {
		panic(err)
	}

	// Optionally check values in redis...
	if got, err := s.Get("foo"); err != nil || got != "bar" {
		t.Error("'foo' has the wrong value")
	}
	// ... or use a helper for that:
	s.CheckGet(t, "foo", "bar")

	// TTL and expiration:
	s.Set("foo", "bar")
	s.SetTTL("foo", 10*time.Second)
	s.FastForward(11 * time.Second)
	if s.Exists("foo") {
		t.Fatal("'foo' should not have existed anymore")
	}
}

// func TestFastagServer(t *testing.T) {
// 	fastagserver.Fastagserver()
// }
