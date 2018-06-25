package main

import "testing"

func TestAdd(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}
	t.Run("3+4", func(t *testing.T) {
		got := Add(3, 4)
		want := 7
		assertCorrectMessage(t, got, want)
	})

	t.Run("1+2", func(t *testing.T) {
		got := Add(1, 2)
		want := 3
		assertCorrectMessage(t, got, want)
	})
}
