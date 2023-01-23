package client

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	v1 "github.com/sarafanfm/mtserver/example/internal/hello/v1"
	v2 "github.com/sarafanfm/mtserver/example/internal/hello/v2"
)

const REQUEST_TIMEOUT = 5 * time.Second

type Client struct {
	grpcEndpoint string
	httpEndpoint string
}

func New(grpcEndpoint, httpEndpoint string) *Client {
	return &Client{grpcEndpoint: grpcEndpoint, httpEndpoint: httpEndpoint}
}

// TODO: gRPC-Web requests emulation
func (c *Client) Run() {
	c.grpcV1()
	c.gatewayV1()

	c.grpcV2()
	c.gatewayV2()
}

func (c *Client) grpcV1() {
	client := v1.NewClient(c.grpcEndpoint)
	defer client.Close()

	connectionCtx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	defer cancel()

	ret, err := client.SayHello(connectionCtx, "Alex")
	if err != nil {
		panic(err)
	}

	log.Printf("grpcV1 response: %v", ret)
}

func (c *Client) grpcV2() {
	client := v2.NewClient(c.grpcEndpoint)
	defer client.Close()

	connectionCtx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	defer cancel()

	ret, err := client.SayHello(connectionCtx, "Alex")
	if err != nil {
		panic(err)
	}

	log.Printf("grpcV2 response: %v", ret)
}

func (c *Client) httpParseResponse(resp *http.Response) string {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func (c *Client) gatewayV1() {
	resp, err := http.Get(c.httpEndpoint + "/hello/v1/Alex")
	if err != nil {
		panic(err)
	}

	log.Printf("gatewayV1 response: %v", c.httpParseResponse(resp))
}

func (c *Client) gatewayV2() {
	resp, err := http.Get(c.httpEndpoint + "/hello/v2/Alex")
	if err != nil {
		panic(err)
	}

	log.Printf("gatewayV2 response: %v", c.httpParseResponse(resp))
}
