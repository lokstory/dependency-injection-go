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

// Inject by type
// @DigoSource
var Service3 IHelloService = &HelloService{Text:"Hello, World 3!"}
