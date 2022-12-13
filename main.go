package main

import (
	configmqtt "rishabhjha12-hub/configMqtt"
	config "rishabhjha12-hub/configRedis"
	configSentry "rishabhjha12-hub/configSentry"
	"rishabhjha12-hub/constants"
)

func main() {
	//Condition checking and initialisation
	configSentry.SentryConfig()
	//Get the gate or init
	//setting enviornment data
	constants.SetEnv()
	//connect to redis
	config.RedisConfig()
	//connect mqtt

	configmqtt.MqttConfig()

	//hit fastag server

	// if helper.Code == constants.Parkzap_Msg_Code {
	// 	fastagserver.Fastagserver()

	// }
}
