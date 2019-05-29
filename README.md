# dependency-injection-go
Say goodbye to "import cycle not allowed"

## Before use 

#### Wonderful package structure

[Trying Clean Architecture on Golang - Iman Tumorang](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)

## Expose variable by upper camel case
    var IHelloService HelloService = &HelloService{}

## Inject dependency by annotation
    // @Inject
    var IHelloService helloService
