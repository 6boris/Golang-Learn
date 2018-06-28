package main

import "testing"

func TestHasPrefix(t *testing.T) {
	str1 := "Hello Word"
	str2 := "He"
	got := HasPrefix(str1, str2)
	want := true

	if got != want {
		t.Errorf("got: '%t' want: '%t'", got, want)
	}
}

func TestHasSuffix(t *testing.T) {
	str1 := "Hello Word"
	str2 := "He"
	got := HasPrefix(str1, str2)
	want := true

	if got != want {
		t.Errorf("got: '%t' want: '%t'", got, want)
	}
}
