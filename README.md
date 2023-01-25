# MultiTransport Server

gRPC + [gRPC-Web](https://github.com/improbable-eng/grpc-web) + [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) + HTTP(S) hybrid server

### GoLang 1.18+ is required

```go
package main

import (
    "github.com/sarafanfm/mtserver"
)

func main() {
    server := mtserver.New()

    // Add endpoints, register gRPC services

    server.Run()
}
```

## Endpoint

Each endpoint is a gRPC or HTTP server, or both. You can specify positive port for enable transport or non-positive for disable it.

```go
http := server.AddEndpoint("http", &mtserver.EndpointOpts{PORT_HTTP: 80})
grpc := server.AddEndpoint("grpc", &mtserver.EndpointOpts{PORT_GRPC: 50021})
hybrid := server.AddEndpoint("hybrid", &mtserver.EndpointOpts{PORT_GRPC: 50021, PORT_HTTP: 80})
```

Also EndpointOpts can:

- `AllowCORS` - enable CORS
- `CORS_Wrapper` - custom CORS func. See default used func [here](https://github.com/sarafanfm/mtserver/blob/main/endpoint.go#L119-L133)
- `UnaryInterceptors` - gRPC unary Middlewares. See [Go gRPC Middleware](https://github.com/grpc-ecosystem/go-grpc-middleware) for more info
- `StreamInterceptors` - gRPC stream Middlewares. See [Go gRPC Middleware](https://github.com/grpc-ecosystem/go-grpc-middleware) for more info
- `GrpcCredentials` - gRPC credentials. `insecure.NewCredentials()` will be used by default
- `TlsDomains` - List of domains for auto-certificates (HTTPS). See [autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert)

Full EndpointOpts declaration:

```go
&mtserver.EndpointOpts{
	PORT_GRPC          int
	PORT_HTTP          int
	AllowCORS          bool
	CORS_Wrapper       func(h http.HandlerFunc) http.HandlerFunc
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	GrpcCredentials    credentials.TransportCredentials
	TlsDomains         []string
}
```

## Register HTTP server

Any HTTP-enabled endpoint have a `Mux`. So you can register any HTTP handler over it.

```go
package main

import (
	"net/http"

	"github.com/sarafanfm/mtserver"
)

func main() {
	server := mtserver.New()
	ep := server.AddEndpoint("http", &mtserver.EndpointOpts{PORT_HTTP: 80})
	root := http.Dir("./static")
	fs := http.FileServer(root)
	ep.Mux.Handle("/", fs)
	server.Run()
}
```

More complex example you can see [here](https://github.com/sarafanfm/mtserver/blob/main/example/internal/static/server.go)

## Register gRPC service

If you have [configured](https://github.com/sarafanfm/mtserver/blob/main/example/protoc.sh#L7-L10) your `protoc` correctly, you should have some interfaces and methods for each `Service` defined in your `*.proto` after compiling proto-files.
Among other things, there should be an `XServer` interface and `RegisterXServer` method where `X` is your service name. Now you can register your service in gRPC endpoint:

```go

type X struct {
	XServer
}

func NewX() *X {
	return &X{}
}

server := mtserver.New()
grpc := server.AddEndpoint("grpc", &mtserver.EndpointOpts{PORT_GRPC: 50021})
mtserver.RegisterGrpcService(
	grpc,
	&mtserver.GrpcService[XServer]{
		Service: NewX(),
		RegisterOnGRPC: RegisterXServer,
	},
)

server.Run()
```

The only thing left for you to do is to implement the methods of your service in `X`.

### gRPC-Web

Will be automatically enabled if HTTP is enabled in your endpoint. So your web-clients must search gRPC server on HTTP port.

### gRPC-Gateway

First, you need to install and configure [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) correctly.
From the example above, your `protoc` compiler must generate one more method - `RegisterXHandlerFromEndpoint`.
Now you can enable handling `google.api.http` proto annotations by changing gRPC service registration with following:

```go
hybrid := server.AddEndpoint("grpc", &mtserver.EndpointOpts{PORT_GRPC: 50021, PORT_HTTP: 80})
mtserver.RegisterGrpcService(
	hybrid,
	&mtserver.GrpcService[XServer]{
		Service: NewX(),
		RegisterOnGRPC: RegisterXServer,
		RegisterOnGateway: RegisterXHandlerFromEndpoint, // <- here. Don't forget to enable HTTP
	},
)
```

Now your server can handle HTTP requests defined in `google.api.http` proto annotations.

## StreamMap

Sometimes we need to receive events from the server. For example: user changed his profile. But we must remember that many other users can subscribe to one event.
For these cases we have included this struct.

Imagine that we have two methods in the service: unary `SaveProfile` and server-side stream `ListenProfile`.
As we know, `protoc` compiler will generate `Svc_ListenProfileServer` struct which will be "response stream" for ListenProfile method.
So if we have integer key as user id then we can do something like this:

```go
type Svc struct {
	// generics is: map key (user id is int), stream type, payload type
	notifyStreams *mtserver.StreamMap[int, Svc_ListenProfileServer, *SomeProfile]
}

func NewSvc() *Svc {
	return &Svc{notifyStreams: mtserver.NewStreamMap[int, Svc_ListenProfileServer, *SomeProfile]()}
}

func (s *Svc) SaveProfile(context context.Context, in *SomeProfile) (*SomeProfile, error) {
	var result *SomeProfile := SomeBusinessLogicSaveMethod()
	s.notifyStreams.Send(result.UserID, result) // <- notify listeners
	return result, nil
}

func (s *Svc) ListenProfile(in *SomeIntegerMessage, stream Svc_ListenProfileServer) error {
	return s.notifyStreams.Add(in.UserID, stream) // <- add listener
}
```

that's all. All disconnected/died streams will be removed automatically. All connected streams will be notified when profile will be changed.

See example [here](https://github.com/sarafanfm/mtserver/blob/main/example/internal/hello/v2/server.go)

## Environment variables

MTServer will try to preload env-formatted (comma separated) files provided in environment variable `ENV_FILE`.

```bash
ENV_FILE=/etc/config.env,./.env go run .
```

See [godotenv](https://github.com/joho/godotenv) for more info.

Also you can use `mtserver.RequiredEnv(string) string` func that will `log.Fatal` if no required env var present.

```go
port := mtserver.RequiredEnv("HTTP_PORT")
```