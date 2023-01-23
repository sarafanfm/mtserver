package v1

import (
	"context"

	"github.com/sarafanfm/mtserver/example/api/hello"
	"github.com/sarafanfm/mtserver/example/api/hello/v1"
)

type Server struct {
	hello.V1Server

	service *Service
}

func NewServer() *Server {
	return &Server{
		service: NewService(),
	}
}

func (s *Server) SayHello(context context.Context, req *v1.Request) (*v1.Response, error) {
	return &v1.Response{Value: s.service.SayHello(req.Value)}, nil
}