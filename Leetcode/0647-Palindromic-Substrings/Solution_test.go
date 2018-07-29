package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := "abc"
		got := countSubstrings(data)
		want := 3
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		data := "aaa"
		got := countSubstrings(data)
		want := 6
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
	t.Run("Test-3", func(t *testing.T) {
		data := "aabaa"
		got := countSubstrings(data)
		want := 6
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
}
