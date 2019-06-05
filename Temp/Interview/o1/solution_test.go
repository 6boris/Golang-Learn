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
		inputs int
		expect int
	}{
		{"TestCase", 1, 3},
		{"TestCase", 2, 3},
		{"TestCase", 3, 3},
		{"TestCase", 4, 3},
		{"TestCase", 5, 3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+"-"+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}
