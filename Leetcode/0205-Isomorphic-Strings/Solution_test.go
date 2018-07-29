package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data1 := "egg"
		data2 := "add"
		got := isIsomorphic(data1, data2)
		want := true
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
	t.Run("Test-2", func(t *testing.T) {
		data1 := "foo"
		data2 := "bar"
		got := isIsomorphic(data1, data2)
		want := false
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
	t.Run("Test-3", func(t *testing.T) {
		data1 := "paper"
		data2 := "title"
		got := isIsomorphic(data1, data2)
		want := true
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

}
