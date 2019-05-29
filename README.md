# dependency-injection-go
Say goodbye to "import cycle not allowed"

## expose variable by upper camel case
    var IHelloService HelloService = &HelloService{}

## Inject dependency by annotation
    // @Inject
    var IHelloService helloService
