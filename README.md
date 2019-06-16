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
* [Installation](#installation)
* [Usage](#usage)
    * [Generation](#generation)
    * [Injecting](#injecting)
* [Annotations](#annotations)
    * [DigoSource](#digosource)
        * [Set sources by type](#set-sources-by-type)
        * [Set sources by name](#set-sources-by-name)
    * [DigoInject](#digoinject)
        * [Inject dependencies by type](#inject-dependencies-by-type)
        * [Inject dependencies by name](#inject-dependencies-by-name)        

-------------------------

## Before use

Recommend putting interfaces and models in separated packages without business logic.

#### Wonderful package structure

[Trying Clean Architecture on Golang - Iman Tumorang](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)


## Installation

```bash
go get -u github.com/lokstory/digo
```

## Usage

### Generation

It will generate below files:
* digo package in the root path of project
* digo.go in every packages which needs to inject dependencies

#### Linux

```bash
go run $GOPATH/src/github.com/lokstory/digo/main.go -path=PROJECT_ROOT_PATH
```

#### Windows

```cmd
go run %GOPATH%\src\github.com\lokstory\digo\main.go -path=PROJECT_ROOT_PATH
```

### Injecting

```go
package main

import (
    "./digo"
)

func main() {
    // Inject dependencies
    digo.Start()
}
```

## Annotations

### @DigoSource

* Expose variables by upper camel case
* Keys must be unique
* Variables must be initialized before call digo.Start()

#### Set sources by type

The source key will be IHelloService.

```go
    // @DigoSource
    // This generator allows to write comments between annotations and variables.
    // Annotations will be ignored if contains empty lines.
    var HelloService IHelloService = &HelloService{}
```

#### Set sources by name

The source key will be HelloService.

```go
    // @DigoSource(HelloService)
    var HelloService IHelloService = &HelloService{}
```    
    
    
### @DigoInject

When using this annotation,
it will create digo.go in the same package to inject the dependency.


#### Inject dependencies by type

```go
    // @DigoInject
    var helloService IHelloService
```

#### Inject dependencies by name

```go
    // @DigoInject(HelloService)
    var helloService IHelloService
```    
