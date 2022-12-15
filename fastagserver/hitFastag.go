package fastagserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	//"io/ioutil"
	"log"
	"net/http"

	"rishabhjha12-hub/constants"
	//"rishabhjha12-hub/helper"
	"rishabhjha12-hub/utils"
	//"github.com/getsentry/sentry-go"
	//"github.com/getsentry/sentry-go"
)

type myStruct struct {
	Epc    string `json:"epc_id"`
	Plzaza string `json:"plaza_id"`
}

var CheckedList []string
var PendingList []string

func Fastagserver(Epckey string, PlazaKey string) {

	//for testing
	//time.Sleep(5 * time.Minute)
	//url := constants.Fastag_Status_Url
	//for testing
	url := "https://uat.fastag.ai/api/v2/icd2.5/fastag_status/"
	method := "POST"
	// Epckey := helper.Epc_ID
	// PlazaKey := helper.Plaza_ID

	//references:https://stackoverflow.com/questions/66842959/submit-variable-in-payload-of-golang-http-newrequest
	myData := myStruct{
		Epc:    Epckey,
		Plzaza: PlazaKey,
	}
	myBytes, _ := json.Marshal(myData)
	payload := bytes.NewBuffer(myBytes)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		//sentry.CaptureException(err)
		return
	}
	req.Header.Add("Authorization", "token 40103c4819c39f878fc68ed8af2191e07d07dca6")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//sentry.CaptureException(err)
		return
	}
	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err)

		return
	}
	fmt.Println(string(body))

	utils.SET_FROM_REDIS(Epckey, string(body), constants.Redis_time_out)
	CheckedList = append(CheckedList, Epckey)

	//utils.GET_FROM_REDIS("fastag_status")

}
