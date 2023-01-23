package util

import (
	"strconv"

	"github.com/sarafanfm/mtserver"
)

func GetHybridPorts() (int, int) {
	grpcPort, err := strconv.Atoi(mtserver.RequiredEnv("PORT_GRPC"))
	if err != nil {
		panic(err)
	}
	httpPort, err := strconv.Atoi(mtserver.RequiredEnv("PORT_HTTP"))
	if err != nil {
		panic(err)
	}
	return grpcPort, httpPort
}