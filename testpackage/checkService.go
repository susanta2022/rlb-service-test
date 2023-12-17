package testpackage

import (
	"fmt"

	mysql_services "github.com/susanta2022/rlb-go-service/services"
)

func CheckTestService() {

	mysql_services.PrintConnectionStatus()

	fmt.Println("response from rlb-service-test >>>>")
}
