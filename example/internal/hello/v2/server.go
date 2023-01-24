package v2

import (
	"context"

	"github.com/sarafanfm/mtserver"
	"github.com/sarafanfm/mtserver/example/api/hello"
	"github.com/sarafanfm/mtserver/example/api/hello/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	hello.V2Server

	service *Service
	notifyStreams *mtserver.StreamMap[string, hello.V2_NotifyHelloServer]
}

func NewServer() *Server {
	return &Server{
		service: NewService(),
		notifyStreams: mtserver.NewStreamMap[string, hello.V2_NotifyHelloServer](),
	}
}

func (s *Server) SayHello(context context.Context, req *v2.Request) (*v2.Response, error) {
	resp := &v2.Response{Val: s.service.SayHello(req.Val)}
	s.notifyStreams.Send("test", resp)
	return resp, nil
}

func (s *Server) NotifyHello(_ *emptypb.Empty, stream hello.V2_NotifyHelloServer) error {
	return s.notifyStreams.Add("test", stream)
}