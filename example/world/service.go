package world

import (
	"../hello"
)

type IService interface {
}

type Service struct {
	IService
}

// @DigoInject
var helloService hello.IHelloService = nil

func Hello() {
	helloService.Hi()
}
