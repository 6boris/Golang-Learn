package Solution

import (
	"fmt"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{73, 74, 75, 71, 69, 72, 76, 73}
		got := dailyTemperatures(data)
		//want := []int{1, 1, 4, 2, 1, 1, 0, 0}

		fmt.Println(got)

	})

}
