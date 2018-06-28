package main

import "testing"

func TestReverseString(t *testing.T) {

	t.Run("Reverse String Test:", func(t *testing.T) {
		got := ReverseString("DEMO")
		want := "OMED"
		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	})

	t.Run("Reverse String Test:", func(t *testing.T) {
		got := ReverseString("ELYK")
		want := "KYLE"
		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	})

	t.Run("Reverse String Test:", func(t *testing.T) {
		got := ReverseString("JIA")
		want := "AIJ"
		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	})

	t.Run("Reverse String Test:", func(t *testing.T) {
		got := ReverseString("JIA")
		want := "AIJ"
		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	})

	t.Run("Reverse String Test:", func(t *testing.T) {
		got := ReverseString("JIA")
		want := "AIJ"
		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	})
}
