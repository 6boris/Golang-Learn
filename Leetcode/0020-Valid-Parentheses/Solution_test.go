package Solution

import (
	"testing"
)

func TestIsValid(t *testing.T) {

	t.Run("--()", func(t *testing.T) {
		got := IsValid("()")
		want := true
		if got != want {
			t.Errorf("WANT:", want, " GOT:", got)
		}
	})

	t.Run("--()[]{}", func(t *testing.T) {
		got := IsValid("()[]{}")
		want := true
		if got != want {
			t.Errorf("WANT:", want, " GOT:", got)
		}
	})

	t.Run("--(]", func(t *testing.T) {
		got := IsValid("(]")
		want := false
		if got != want {
			t.Errorf("WANT:", want, " GOT:", got)
		}
	})

	t.Run("--([)]", func(t *testing.T) {
		got := IsValid("([)]")
		want := false
		if got != want {
			t.Errorf("WANT:", want, " GOT:", got)
		}
	})

	t.Run("--{[]}", func(t *testing.T) {
		got := IsValid("{[]}")
		want := true
		if got != want {
			t.Errorf("WANT:", want, " GOT:", got)
		}
	})
}
