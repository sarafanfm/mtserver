package main

import (
	"strconv"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sarafanfm/mtserver"
	"google.golang.org/grpc"

	"github.com/sarafanfm/mtserver/example/util"
	api "github.com/sarafanfm/mtserver/example/api/hello"

	helloV1 "github.com/sarafanfm/mtserver/example/internal/hello/v1"
	helloV2 "github.com/sarafanfm/mtserver/example/internal/hello/v2"

	static "github.com/sarafanfm/mtserver/example/internal/static"
)

func main() {
	server := mtserver.New()

	/////////////////////////// HYBRID /////////////////////////////////////

	grpcPort, httpPort := util.GetHybridPorts()

	hybrid := server.AddEndpoint("hybrid", &mtserver.EndpointOpts{
		PORT_GRPC: grpcPort,
		PORT_HTTP: httpPort,
		//AllowCORS: true,
		UnaryInterceptors: []grpc.UnaryServerInterceptor{
			grpc_recovery.UnaryServerInterceptor(),
		},
		StreamInterceptors: []grpc.StreamServerInterceptor{
			grpc_recovery.StreamServerInterceptor(),
		},
		//TlsDomains: []string{"example.com"}, // auto https with letsencrypt
	})

	mtserver.RegisterGrpcService(
		hybrid,
		&mtserver.GrpcService[api.V1Server]{
			Service:           helloV1.NewServer(),
			RegisterOnGRPC:    api.RegisterV1Server,
			RegisterOnGateway: api.RegisterV1HandlerFromEndpoint,
		},
	)

	mtserver.RegisterGrpcService(
		hybrid,
		&mtserver.GrpcService[api.V2Server]{
			Service:           helloV2.NewServer(),
			RegisterOnGRPC:    api.RegisterV2Server,
			RegisterOnGateway: api.RegisterV2HandlerFromEndpoint,
		},
	)

	// Add HTTP static server with root URI NOT "/" !!!
	static.NewServer(hybrid.Mux, "/static/")

	//////////////////////////// HTTP //////////////////////////////////////

	httpPort, err := strconv.Atoi(mtserver.RequiredEnv("PORT_HTTP_STATIC"))
	if err != nil {
		panic(err)
	}
	http := server.AddEndpoint("http", &mtserver.EndpointOpts{PORT_GRPC: -1, PORT_HTTP: httpPort})
	static.NewServer(http.Mux, "/")

	////////////////////////////////////////////////////////////////////////

	server.Run()
}
