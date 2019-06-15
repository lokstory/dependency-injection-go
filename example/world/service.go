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
// Inject dependency by type
var helloService hello.IHelloService

func Hello() {
	helloService.Hi()
}
