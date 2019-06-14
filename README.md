# Digo

<p align="center">
  <a href="https://goreportcard.com/report/github.com/lokstory/digo"><img src="https://goreportcard.com/badge/github.com/lokstory/digo"></a>
  <a href="https://godoc.org/github.com/lokstory/digo"><img src="https://godoc.org/github.com/lokstory/digo?status.svg" alt="GoDoc"></a>
</p>

*In Taiwanese, the mean of "Digo" is "Pig Brother".*

-------------------------

### Say goodbye to "import cycle not allowed"

-------------------------

* [Before use](#before-use)
* [Annotations](#annotations)
    * [DigoSource](#digosource)
        * [Inject the dependency by type](#inject-the-dependency-by-type)
        * [Inject the dependency by name](#inject-the-dependency-by-name)
    * [DigoInject](#digoinject)

-------------------------

## Before use 

#### Wonderful package structure

[Trying Clean Architecture on Golang - Iman Tumorang](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)

## Annotations

### @DigoSource

* Expose variable by upper camel case
* The source key must be unique


#### Setting the source by type

The source key will be IHelloService

```go
    // @DigoSource
    var HelloService IHelloService = &HelloService{}
```

#### Setting the source by name

The source key will be HelloService

```go
    // @DigoSource(HelloService)
    var HelloService IHelloService = &HelloService{}
```    
    
    
### @DigoInject

#### Injecting the dependency by type

```go
    // @DigoSource
    var helloService IHelloService
```

#### Injecting the dependency by name

```go
    // @DigoSource(HelloService)
    var helloService IHelloService
```    
