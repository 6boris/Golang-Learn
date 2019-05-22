package o1

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect int
	}{
		{"TestCase", []int{1, 2}, 3},
		{"TestCase", []int{2, 2}, 4},
		{"TestCase", []int{3, 2}, 5},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+"-"+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs[0], c.inputs[1])
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}
