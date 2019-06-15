package test

import (
	// It will be generated after building
	"../digo"
	"../world"
	"testing"
)

// TestDigo
func TestDigo(t *testing.T) {
	digo.Start()
	world.Hello()
}
