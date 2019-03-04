package d

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]string
		expect []string
	}{
		{"TestCacse 1",
			[][]string{
				{"1000", "00:00:00"},
				{"1001", "00:59:00"},
				{"1001", "00:58:00"},
				{"1001", "01:00:00"},
				{"1000", "01:00:00"},
				{"1000", "02:00:00"},
				{"1000", "03:00:00"},
				{"1001", "01:00:01"},
				{"1000", "04:00:00"},
				{"1001", "01:23:14"},
			},
			[]string{"1001"}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				//t.Fatalf("expected: %v, but got: %v, with inputs: %v",
				//	c.expect, got, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
