package main

import (
	"runtime"
	"testing"
)

func TestSlieve(t *testing.T) {
	t.Run("GOMAXPROCS == 1", func(t *testing.T) {
		runtime.GOMAXPROCS(1)
		Slieve()
	})
	t.Run("GOMAXPROCS == 4", func(t *testing.T) {
		runtime.GOMAXPROCS(50)
		Slieve()
	})
}
