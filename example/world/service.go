package world

import (
	"../hello"
)

type IService interface {
}

type Service struct {
	IService
}

// @DigoInject(HelloService)
var helloService hello.IHelloService
