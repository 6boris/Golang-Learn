package v2

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestStackMain(t *testing.T) {
	t.Run("PUSH int", func(t *testing.T) {

		s := Stack{}
		for i := 0; i < 1000; i++ {
			s.Push(i)
		}
		if s.Top() != 999 {
			t.Error("type error")
		}

		if s.Pop() != 999 {
			t.Error("errr of pop")
		}
		if s.Pop() != 998 {
			t.Error("errr of pop")
		}
	})
	t.Run("PUSH time", func(t *testing.T) {
		s := Stack{}
		for i := 0; i < 100000; i++ {
			s.Push(time.Now())
		}
	})

	t.Run("PUSH strings", func(t *testing.T) {
		s := Stack{}
		for i := 0; i < 100000; i++ {
			s.Push(string("asd"))
		}

		if s.Top() != "a s d" {
			t.Error("stringc error")
		}
		a := s.Top().(string)
		res := strings.Split(a, " ")
		fmt.Println(res)
	})
}
