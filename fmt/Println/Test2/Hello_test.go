package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("English:", func(t *testing.T) {
		got := Hello("english", "Kyle")

		want := "Hello, Kyle"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Spanish:", func(t *testing.T) {
		got := Hello("spanish", "Kyle")
		want := "Spanish, Kyle"
		assertCorrectMessage(t, got, want)
	})

	t.Run("French:", func(t *testing.T) {

	})

}
