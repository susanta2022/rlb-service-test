package testpackage

import (
	"fmt"

	mysql_services "github.com/susanta2022/rlb-go-service/services"
)

func main() {
	mysql_services.PrintConnectionStatus()
	fmt.Print("susanta")
}
