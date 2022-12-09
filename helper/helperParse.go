package helper

import (
	"encoding/json"
	"fmt"
	"log"

	"rishabhjha12-hub/constants"
	//"rishabhjha12-hub/fastagserver"
	"rishabhjha12-hub/utils"
)

var Epc_ID string
var Plaza_ID string
var Code interface{}

func ParseHelper(myslice []byte) {
	// Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(myslice), &result)
	fmt.Println("message type code", result["message_type_code"])

	code := result["message_type_code"]
	Code = code

	//distributing the code according to message codes
	if code == constants.Parkzap_Msg_Code {
		log.Println("topic : Parkzap")
		log.Println("epc id: ", result["message_body"].(map[string]interface{})["epc_id"])
		log.Println("plaza id: ", result["message_body"].(map[string]interface{})["plaza_id"])
		message_body2 := fmt.Sprint(result["message_body"].(map[string]interface{})["epc_id"])
		message_body11 := fmt.Sprint(result["message_body"].(map[string]interface{})["plaza_id"])
		utils.SET_FROM_REDIS(constants.Epc_key, message_body2, constants.Redis_time_out)
		utils.GET_FROM_REDIS(constants.Epc_key)
		Epc_ID = message_body2
		Plaza_ID = message_body11
		//fastagserver.Fastagserver()

	} else if code == constants.Jetson_Msg_Code {
		log.Println("topic : jetson")
		log.Println("anpr: ", result["message_body"].(map[string]interface{})["data_corrected"])
		message_body1 := fmt.Sprint(result["message_body"].(map[string]interface{})["data_corrected"])
		//setting data to redis

		utils.SET_FROM_REDIS(constants.Anpr_key, message_body1, constants.Redis_time_out)
		//getting data from redis
		utils.GET_FROM_REDIS(constants.Anpr_key)

	} else if code == constants.Parkzap_loop_code {
		log.Println("topic : loop")

		isLoopHigh := result["message_body"].(map[string]interface{})["signal_state"]
		log.Println(isLoopHigh)
		if isLoopHigh == "low" {
			utils.SET_FROM_REDIS(constants.Loop_key, "false", constants.Redis_time_out)

		} else {
			utils.SET_FROM_REDIS(constants.Loop_key, "true", constants.Redis_time_out)

		}
		utils.GET_FROM_REDIS(constants.Loop_key)

	} else {
		log.Println("topic : other")
	}

}
