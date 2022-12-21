package main

import (
	//"encoding/json"
	"fmt"
	configmqtt "rishabhjha12-hub/configMqtt"
	config "rishabhjha12-hub/configRedis"
	configsentry "rishabhjha12-hub/configSentry"
	"rishabhjha12-hub/constants"
	"rishabhjha12-hub/fastagserver"
	"rishabhjha12-hub/helper"
	"rishabhjha12-hub/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/assert"
)

// this will test the redis configuration
func TestRedis(t *testing.T) {
	config.RedisConfig()
}

// this functon will try to set the value of keyss=50 and then check if the value of the keyss is nil or not nil if it will be nil then it will throw an error
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

// test sentry configuration
func TestConfigSentry(t *testing.T) {
	configsentry.SentryConfig()
}

// it runs the function which listen to mqtt and the response
func TestUnmarshalAndCheckResponse(t *testing.T) {
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
		assert.Equal(t, val, constants.Parkzap_Msg_Code, "should be equal")

		fmt.Println("*******======", val)
	})
}

// this function will test configuration of mqtt
func TestMqtt(t *testing.T) {
	configmqtt.MqttConfig()
}

// this function will hit fastag server with epc key and plaza id and compare the actual response with the response coming from the function
func TestFastagServer(t *testing.T) {
	Epckey := "34161FA8203288AC05FF4500"
	Plaza_ID := "15"
	val := fastagserver.Fastagserver(Epckey, Plaza_ID)
	myres := `{"response_id": 201, "bank_id": null, "epc_id": "34161FA8203288AC05FF4500", "t_id": "", "vehicle_number": "", "blacklist_status": true, "chassis_number": "", "vehicle_class": "", "is_commercial": "F", "status_code": 400, "tag_history_vehicle": null}`
	assert.Equal(t, string(myres), val, "should be equal")
}
