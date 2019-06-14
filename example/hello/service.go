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

// @DigoSource(HelloService)
var Service IHelloService = &HelloService{Text:"Hello, World!"}

// @DigoSource
var Service2 IHelloService = &HelloService{Text:"Hello, World 2!"}
