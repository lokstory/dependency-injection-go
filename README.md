<p align="center"><img src="https://raw.githubusercontent.com/lokstory/digo/master/digo.png"></p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/lokstory/digo"><img src="https://goreportcard.com/badge/github.com/lokstory/digo"></a>
  <a href="https://godoc.org/github.com/lokstory/digo"><img src="https://godoc.org/github.com/lokstory/digo?status.svg" alt="GoDoc"></a>
</p>

# Digo

*In Taiwanese, the mean of "Digo" is "Pig Brother".*

### Say goodbye to "import cycle not allowed"!

-------------------------

* [Before use](#before-use)
* [Annotations](#annotations)
    * [DigoSource](#digosource)
        * [Setting the source by type](#setting-the-source-by-type)
        * [Setting the source by name](#setting-the-source-by-name)
    * [DigoInject](#digoinject)
        * [Inject the dependency by type](#injecting-the-dependency-by-type)
        * [Inject the dependency by name](#injecting-the-dependency-by-name)

-------------------------

## Before use 

#### Wonderful package structure

[Trying Clean Architecture on Golang - Iman Tumorang](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)

## Annotations

### @DigoSource

* Expose variable by upper camel case
* The source key must be unique
* The variable must be initialized before call digo.Start()


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

When using this annotation,
it will create digo.go in the same package to inject the dependency.


#### Injecting the dependency by type

```go
    // @DigoInject
    var helloService IHelloService
```

#### Injecting the dependency by name

```go
    // @DigoInject(HelloService)
    var helloService IHelloService
```    
