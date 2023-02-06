package mtserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type EndpointOpts struct {
	PORT_GRPC          int
	PORT_HTTP          int
	AllowCORS          bool
	CORS_Wrapper       func(h http.HandlerFunc) http.HandlerFunc
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	GrpcCredentials    credentials.TransportCredentials
	TlsDomains         []string

	OnStart         Callback
	OnStartError    CallbackError
	OnShutdown      Callback
	OnForceShutdown Callback
}

type GrpcService[T any] struct {
	Service           T
	RegisterOnGRPC    func(s grpc.ServiceRegistrar, srv T)
	RegisterOnGateway func(context.Context, *gwruntime.ServeMux, string, []grpc.DialOption) (err error)
}

type Endpoint struct {
	grpc    *grpc.Server
	grpcWeb *grpcweb.WrappedGrpcServer
	gateway *gwruntime.ServeMux
	http    *http.Server
	Mux     *http.ServeMux
	options *EndpointOpts
}

func newEndpoint(name string, opts *EndpointOpts) *Endpoint {
	ret := &Endpoint{}
	if opts == nil {
		opts = &EndpointOpts{}
	}

	if opts.PORT_GRPC <= 0 && opts.PORT_HTTP <= 0 {
		panic("cannot create Endpoint if gRPC and HTTP both disabled")
	}

	ret.options = opts

	if opts.PORT_GRPC > 0 {
		unaryInterceptor := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer((opts.UnaryInterceptors)...))
		streamInterceptor := grpc.StreamInterceptor(grpc_middleware.ChainStreamServer((opts.StreamInterceptors)...))

		ret.grpc = grpc.NewServer(unaryInterceptor, streamInterceptor)

		if opts.PORT_HTTP > 0 {
			ret.grpcWeb = grpcweb.WrapServer(ret.grpc)
			ret.gateway = gwruntime.NewServeMux()
		}
	}

	if opts.PORT_HTTP > 0 {
		ret.Mux = http.NewServeMux()

		if ret.gateway != nil {
			ret.Mux.Handle("/", ret.gateway)
		}

		handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			if ret.grpcWeb != nil {
				if ret.grpcWeb.IsGrpcWebRequest(req) {
					ret.grpcWeb.ServeHTTP(resp, req)
					return
				}
			}

			ret.Mux.ServeHTTP(resp, req)
		})

		if opts.AllowCORS {
			if opts.CORS_Wrapper != nil {
				handler = opts.CORS_Wrapper(handler)
			} else {
				handler = allowCORS(handler)
			}
		}

		ret.http = &http.Server{
			Addr:    fmt.Sprintf(":%d", opts.PORT_HTTP),
			Handler: handler,
		}
	}

	return ret
}

func RegisterGrpcService[T any](endpoint *Endpoint, svc *GrpcService[T]) {
	if endpoint.grpc == nil {
		panic("cannot register gRPC service on non-gRPC endpoint")
	}
	svc.RegisterOnGRPC(endpoint.grpc, svc.Service)
	if endpoint.gateway != nil && svc.RegisterOnGateway != nil {
		creds := endpoint.options.GrpcCredentials
		if creds == nil {
			creds = insecure.NewCredentials()
		}
		svc.RegisterOnGateway(context.Background(), endpoint.gateway, fmt.Sprintf("127.0.0.1:%d", endpoint.options.PORT_GRPC), []grpc.DialOption{grpc.WithTransportCredentials(creds)})
	}
}

func allowCORS(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin) // TODO: allowed origins
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				headers := []string{"Content-Type", "Accept", "Authorization", "X-User-Agent", "X-gRPC-Web"}
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
				methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
