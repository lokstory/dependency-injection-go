# digo
Say goodbye to "import cycle not allowed"

## Before use 

#### Wonderful package structure

[Trying Clean Architecture on Golang - Iman Tumorang](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)

#### Expose variable by upper camel case, and set the source name by annotation
    // @DigoSource(HelloService)
    var HelloService IHelloService = &HelloService{}

#### Inject dependency by annotation
    // @DigoInject
    var IHelloService helloService
