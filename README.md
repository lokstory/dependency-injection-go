
[![Go Report Card](https://goreportcard.com/badge/github.com/lokstory/digo)](https://goreportcard.com/report/github.com/lokstory/digo)
[![GoDoc](https://godoc.org/github.com/lokstory/digo?status.svg)](https://godoc.org/github.com/lokstory/digo)

# digo
Say goodbye to "import cycle not allowed"

## Before use 

#### Wonderful package structure

[Trying Clean Architecture on Golang - Iman Tumorang](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)

## Source

##### Expose variable by upper camel case
##### The name or type must be a unique key

#### Inject by type
##### The source key will be IHelloService
    // @DigoSource
    var HelloService IHelloService = &HelloService{} 

#### Inject by name
##### The source key will be HelloService  
    // @DigoSource(HelloService)
    var HelloService IHelloService = &HelloService{} 
    
    
## Target
#### Inject dependency by annotation
    // @DigoInject(HelloService)
    var helloService IHelloService
