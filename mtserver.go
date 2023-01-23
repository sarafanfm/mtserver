package mtserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const SHUTDOWN_TIMEOUT_SECONDS = 5

type MTServer struct {
	jobs []*Job
	p    map[string]*Endpoint
}

func New() *MTServer {
	m := &MTServer{}
	m.p = make(map[string]*Endpoint)
	return m
}

func (m *MTServer) Run() {
	run(m.jobs...)
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
					log.Fatalf("failed to start gRPC (HTTP/2) server %s: %v", name, err)
					return err
				}

				log.Printf("starting gRPC (HTTP/2) server %s at %s", name, lis.Addr().String())

				return endpoint.grpc.Serve(lis)
			},
			Shutdown: func() {
				log.Printf("shutting down gRPC (HTTP/2) server %s", name)

				go func() {
					time.Sleep(time.Second * SHUTDOWN_TIMEOUT_SECONDS)
					log.Printf("force terminating gRPC (HTTP/2) server %s", name)
					endpoint.grpc.Stop()
				}()

				endpoint.grpc.GracefulStop()
			},
		})
	}
	if endpoint.http != nil {
		m.jobs = append(m.jobs, &Job{
			Run: func() error {
				log.Printf("starting HTTP server %s at %s", name, endpoint.http.Addr)

				err := endpoint.http.ListenAndServe()

				if err == http.ErrServerClosed {
					return nil
				}
				return err
			},
			Shutdown: func() {
				log.Printf("shutting down HTTP server %s", name)

				ctx, cancel := context.WithTimeout(context.Background(), SHUTDOWN_TIMEOUT_SECONDS*time.Second)
				defer cancel()

				endpoint.http.Shutdown(ctx)
			},
		})
	}

	return endpoint
}

func (m *MTServer) GetEndpoint(name string) *Endpoint {
	return m.p[name]
}
