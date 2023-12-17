package testpackage

import (
	"fmt"

	mysql_services "github.com/susanta2022/rlb-go-service/services"
)

func check_test_service() {
	mysql_services.PrintConnectionStatus()
	fmt.Print("response from rlb-service-test")
}
