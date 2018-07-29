package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		s := "abccccdd"
		got := longestPalindrome(s)
		want := 7
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

}
