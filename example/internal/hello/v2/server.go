package v2

import (
	"context"

	"github.com/sarafanfm/mtserver/example/api/hello"
	"github.com/sarafanfm/mtserver/example/api/hello/v2"
)

type Server struct {
	hello.V2Server

	service *Service
}

func NewServer() *Server {
	return &Server{
		service: NewService(),
	}
}

func (s *Server) SayHello(context context.Context, req *v2.Request) (*v2.Response, error) {
	return &v2.Response{Val: s.service.SayHello(req.Val)}, nil
}