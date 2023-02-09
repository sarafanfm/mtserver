package mtserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

const SHUTDOWN_TIMEOUT = 5 * time.Second

type Callback func()
type CallbackError func(error)

type MTServer struct {
	OnJobError CallbackError

	jobs []*Job
	p    map[string]*Endpoint
}

func New() *MTServer {
	m := &MTServer{}
	m.p = make(map[string]*Endpoint)
	return m
}

func (m *MTServer) Run() {
	run(m.OnJobError, m.jobs...)
}

func (m *MTServer) AddJob(j *Job) {
	m.jobs = append(m.jobs, j)
}

func (m *MTServer) AddEndpoint(name string, opts *EndpointOpts) *Endpoint {
	endpoint := newEndpoint(name, opts)

	m.p[name] = endpoint

	if endpoint.grpc != nil {
		m.jobs = append(m.jobs, &Job{
			Run: func() error {
				lis, err := net.Listen("tcp", fmt.Sprintf(":%d", endpoint.options.PORT_GRPC))
				if err != nil {
					if endpoint.options.OnStartError != nil {
						endpoint.options.OnStartError(err)
					}
					return err
				}

				if endpoint.options.OnStart != nil {
					endpoint.options.OnStart()
				}

				return endpoint.grpc.Serve(lis)
			},
			Shutdown: func() {
				if endpoint.options.OnShutdown != nil {
					endpoint.options.OnShutdown()
				}

				time.AfterFunc(SHUTDOWN_TIMEOUT, func() {
					if endpoint.options.OnForceShutdown != nil {
						endpoint.options.OnForceShutdown()
					}
					endpoint.grpc.Stop()
				})

				endpoint.grpc.GracefulStop()
			},
		})
	}
	if endpoint.http != nil {
		m.jobs = append(m.jobs, &Job{
			Run: func() error {
				if endpoint.options.OnStart != nil && endpoint.grpc == nil {
					endpoint.options.OnStart()
				}

				var err error

				if opts.TlsDomains != nil && len(opts.TlsDomains) > 0 {
					err = endpoint.http.Serve(autocert.NewListener(opts.TlsDomains...))
				} else {
					err = endpoint.http.ListenAndServe()
				}

				if err == http.ErrServerClosed {
					return nil
				}
				return err
			},
			Shutdown: func() {
				if endpoint.options.OnShutdown != nil && endpoint.grpc == nil {
					endpoint.options.OnShutdown()
				}

				ctx, cancel := context.WithTimeout(context.Background(), SHUTDOWN_TIMEOUT)
				defer func() {
					if endpoint.options.OnForceShutdown != nil {
						endpoint.options.OnForceShutdown()
					}
					cancel()
				}()

				endpoint.http.Shutdown(ctx)
			},
		})
	}

	return endpoint
}

func (m *MTServer) GetEndpoint(name string) *Endpoint {
	return m.p[name]
}
