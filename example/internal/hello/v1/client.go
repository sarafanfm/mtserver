package v1

import (
	"context"
	"time"

	api "github.com/sarafanfm/mtserver/example/api/hello"
	v1 "github.com/sarafanfm/mtserver/example/api/hello/v1"
	"github.com/sarafanfm/mtserver/example/internal/common"
	"google.golang.org/grpc"
)

type Client struct {
	common.Client

	srv api.V1Client
}

func NewClient(address string) *Client {
	c := &Client{}
	c.CreateConnection(address)
	c.srv = api.NewV1Client(c.Conn)
	return c
}

func (c *Client) SayHello(ctx context.Context, in string, opts ...grpc.CallOption) (*v1.Response, error) {
	connectionCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return c.srv.SayHello(connectionCtx, &v1.Request{Value: in}, opts...)
}
