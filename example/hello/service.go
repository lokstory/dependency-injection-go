package hello

import "log"

type IHelloService interface {
	Hi()
}

type HelloService struct {
	IHelloService
	Text string
}

func (s *HelloService) Hi() {
	log.Println(s.Text)
}

// Inject by key
// @DigoSource(HelloService)
var Service IHelloService = &HelloService{Text:"Hello, World!"}
