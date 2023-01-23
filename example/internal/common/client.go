package common

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Conn *grpc.ClientConn
}

func (c *Client) CreateConnection(address string) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect to gateway %s: %v", address, err)
	}
	c.Conn = conn
}

func (c *Client) Close() {
	c.Conn.Close()
}