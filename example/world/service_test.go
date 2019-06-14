package world

import (
	"log"
	"strings"
	"testing"
)

// @DepSource(HelloService)
// @DepInject(HelloService)

// @Source(HelloService)
// @DigoSource(HelloService)

// #DigoSet(HelloService)
// #DigoGet(HelloService)

// @DigoSet(HelloService)
// @DigoGet(HelloService)

// @DigoSource(HelloService)
// @DigoInject(HelloService)

// @Digo(source="HelloService")

// @Digo({"source":"HelloService"})
// @Digo({"inject":"HelloService"})
func TestSetService(t *testing.T) {
	text := `   \\@Inject(Hello)`
	log.Println(strings.TrimLeft(text, "\\@Inject"))
}
