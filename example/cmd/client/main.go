package main

import (
	"fmt"

	"github.com/sarafanfm/mtserver/example/util"
	"github.com/sarafanfm/mtserver/example/internal/client"
)

func main() {
	grpcPort, httpPort := util.GetHybridPorts()

	grpcAddress := fmt.Sprintf("localhost:%d", grpcPort)
	httpAddress := fmt.Sprintf("http://localhost:%d", httpPort)
	
	worker := client.New(grpcAddress, httpAddress)

	worker.Run()
}
