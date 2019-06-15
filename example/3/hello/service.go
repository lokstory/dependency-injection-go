package hello

import (
	"log"
)

// Test duplicate interface

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
// @DigoSource(HelloService3)
var Service3 IHelloService = &HelloService{Text:"Hello, 3!"}
