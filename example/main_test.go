package example

import (
	"../digo"
	"./world"
	"testing"
)

func TestDigo(t *testing.T) {
	digo.Start()
	world.Hello()
}
