package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := 3
		got := judgeSquareSum(data)
		want := false
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
	t.Run("Test-2", func(t *testing.T) {
		data := 5
		got := judgeSquareSum(data)
		want := true
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
}
