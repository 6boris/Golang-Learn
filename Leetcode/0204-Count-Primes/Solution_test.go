package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := countPrimes(10)
		want := 4
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
			//t.Errorf(string(got))
		}
	})
}
