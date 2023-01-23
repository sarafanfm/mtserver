package v1

type Service struct {}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SayHello(req string) string {
	return "Hello, " + req + "!"
}