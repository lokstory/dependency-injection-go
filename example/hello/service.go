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

// Set the source by type
// @DigoSource
var Service1 IHelloService = &HelloService{Text:"Hello, 1!"}
