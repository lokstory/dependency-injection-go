package world

import (
	dependencyContract "../../dependency/contract"
	"../hello"
	"log"
	"reflect"
)

type IService interface {
}

type Service struct {
	IService
}

// @DigoInject(HelloService)
var helloService hello.IHelloService

func InitDependency(manager dependencyContract.IManager) {
	log.Println("param:", manager.GetPointer("hello"))

	// Unsafe cast
	//helloService = *(*hello.IHelloService)(manager.GetPointer("hello"))

	// Reflection
	ptr := reflect.ValueOf(manager.GetHello()).Elem().Addr()
	reflect.ValueOf(&helloService).Elem().Set(ptr)

	log.Println("hello:", helloService, &helloService, helloService == nil)

	helloService.Hi()
}
