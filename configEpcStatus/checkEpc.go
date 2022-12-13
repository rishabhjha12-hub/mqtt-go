package configepcstatus

import (
	"rishabhjha12-hub/fastagserver"
	"sync"

	"golang.org/x/exp/slices"
	//"rishabhjha12-hub/fastagserver"
)

var Wg sync.WaitGroup

func GetEpc(Epc string, plaza string) {
	var CheckedList []string
	var PendingList []string
	if len(Epc) == 24 {
		if slices.Contains(CheckedList, Epc) {
			//save time stamp

		} else {
			if slices.Contains(PendingList, Epc) {
				//save time stamps

			} else {
				go fastagserver.Fastagserver(Epc, plaza)

			}

		}
	}

}
