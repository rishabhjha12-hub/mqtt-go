package main

import (
	"fmt"
	configmqtt "rishabhjha12-hub/configMqtt"
	config "rishabhjha12-hub/configRedis"
	configSentry "rishabhjha12-hub/configSentry"
	"rishabhjha12-hub/constants"
	"rishabhjha12-hub/fastagserver"
	"rishabhjha12-hub/helper"
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
	fmt.Println("helper code", helper.Code)
	fmt.Println("parkzap msg code", constants.Parkzap_Msg_Code)
	//if helper.Code == constants.Parkzap_Msg_Code {
	fastagserver.Fastagserver()

	//}
}
