package fastagserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rishabhjha12-hub/constants"
	"rishabhjha12-hub/helper"
	"rishabhjha12-hub/utils"
)

type myStruct struct {
	Epc    string `json:"epc_id"`
	Plzaza string `json:"plaza_id"`
}

func Fastagserver() {

	url := "https://uat.fastag.ai/api/v2/icd2.5/fastag_status/"
	method := "POST"
	Epckey := helper.Epc_ID
	PlazaKey := helper.Plaza_ID

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
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "token 40103c4819c39f878fc68ed8af2191e07d07dca6")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	utils.SET_FROM_REDIS("fastag_status", string(body), constants.Redis_time_out)
	//utils.GET_FROM_REDIS("fastag_status")

}
