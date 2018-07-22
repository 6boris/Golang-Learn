package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello Word"
	if got != want {
		t.Errorf("got: '%s' want: '%s'", got, want)
	}
}
