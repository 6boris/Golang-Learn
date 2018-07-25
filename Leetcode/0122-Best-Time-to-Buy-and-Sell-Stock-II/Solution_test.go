package Solution

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	data := []int{7, 1, 5, 3, 6, 4}
	got := maxProfit(data)
	want := 5
	if got != want {
		t.Errorf("GOT:", got, " WANT:", want)
	}
}
