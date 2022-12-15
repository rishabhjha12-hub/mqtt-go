package utils

func ValidateEpc(Epc string) bool {
	if len(Epc) == 24 && Epc[0] == 3 && Epc[1] == 4 && Epc[2] == 1 {
		return true

	} else {
		return false
	}
}
