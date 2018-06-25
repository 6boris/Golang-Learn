package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Say Hello to SomeOne:", func(t *testing.T) {
		got := Hello("Kyle")
		want := "Hello Kyle"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello to SomeOne when empty string supplied:", func(t *testing.T) {
		got := Hello("")
		want := "Hello Kyle"
		assertCorrectMessage(t, got, want)
	})

}
