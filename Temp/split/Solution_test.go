package Solution

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []string
		expect string
	}{
		{"某某访问了您的名片****[]【】, 2018-06-06 00:00:00",
			[]string{"某某访问了您的名片****[]【】, 2018-06-06 00:00:00", ","},
			"某某访问了您的名片****[]【】",
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Split(c.inputs[0], c.inputs[1])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
