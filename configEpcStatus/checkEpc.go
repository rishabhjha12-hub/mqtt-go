package configepcstatus

import (
	"log"
	"rishabhjha12-hub/constants"
	"rishabhjha12-hub/fastagserver"

	"rishabhjha12-hub/utils"
	"sync"
	"time"

	"golang.org/x/exp/slices"
	//"rishabhjha12-hub/fastagserver"
)

var Wg sync.WaitGroup

func GetEpc(Epc string, plaza string) {

	//utils.SET_FROM_REDIS("StatusCheckedList",CheckedList[],constants.Redis_time_out)
	if len(Epc) == 24 {
		if slices.Contains(fastagserver.CheckedList, Epc) {
			//save time stamp
			utils.SET_FROM_REDIS("lEpc", time.Now().String(), constants.Redis_time_out)

		} else {
			if slices.Contains(fastagserver.PendingList, Epc) {
				//save time stamps
				utils.SET_FROM_REDIS("lEpc", time.Now().String(), constants.Redis_time_out)

			} else {

				//for testing
				log.Println("going in")
				go fastagserver.Fastagserver(Epc, plaza)
				//for testing
				log.Println("going out")

			}

		}
	}

}
