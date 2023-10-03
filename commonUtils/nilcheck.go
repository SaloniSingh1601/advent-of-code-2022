package nilcheck

import "log"

func CheckNilErr(err error) {
	if err != nil {
		log.Fatal("Error:",err)
	}
}