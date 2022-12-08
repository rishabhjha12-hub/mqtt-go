package main

import (
	config "rishabhjha12-hub/configRedis"
	configsentry "rishabhjha12-hub/configSentry"
	"rishabhjha12-hub/fastagserver"
	"rishabhjha12-hub/utils"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	config.RedisConfig()

}
func TestGetterSetter(t *testing.T) {
	utils.SET_FROM_REDIS("keyss", "50", 5*time.Second)
	utils.GET_FROM_REDIS("keyss")

}
func TestConfigSentry(t *testing.T) {
	configsentry.SentryConfig()
}

func TestFastagServer(t *testing.T) {
	fastagserver.Fastagserver()
}

// func TestMqtt(t *testing.T) {
// 	configmqtt.MqttConfig()
// }
