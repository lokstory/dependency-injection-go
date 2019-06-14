package hello

// Inject by key
// @DigoSource(HelloService2)
var Service2 IHelloService = &HelloService{Text:"Hello, World 2!"}
