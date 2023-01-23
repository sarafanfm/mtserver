package v2

import "github.com/sarafanfm/mtserver/example/internal/hello/v1"

type Service struct {
	v1.Service
}

func NewService() *Service {
	return &Service{}
}

// SayHello is a method of v1 Service.