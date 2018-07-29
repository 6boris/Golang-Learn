package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		s1 := "anagram"
		s2 := "nagaram"
		got := isAnagram(s1, s2)
		want := true
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		s1 := "rat"
		s2 := "car"
		got := isAnagram(s1, s2)
		want := false
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

}
