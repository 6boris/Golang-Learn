package Solution

import (
	"testing"
)

func TestcheckPossibility(t testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{4, 2, 3}
		got := checkPossibility(data)
		want := true

		if got != want {
			t.Errorf("Test Fail")
			//t.Errorf("GOT:", got, "WANT", want)
		}
	})
	t.Run("Test-2", func(t *testing.T) {
		data := []int{4, 2, 1}
		got := checkPossibility(data)
		want := false

		if got != want {
			t.Errorf("Test Fail")
			//t.Errorf("GOT:", got, "WANT", want)
		}
	})
}
