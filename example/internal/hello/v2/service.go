package v2

import v1 "github.com/sarafanfm/mtserver/example/internal/hello/v1"

type Service struct {
	v1.Service
}

func NewService() *Service {
	return &Service{}
}

// SayHello is a method of v1 Service, equivalent to:
/*
func (s *Service) SayHello(val string) string {
	return s.Service.SayHello(val)
}
*/
