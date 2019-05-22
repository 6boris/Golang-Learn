package main

import (
	"fmt"
	"sort"
)

type S struct {
	v int
}

func main() {
	s := []S{{1}, {3}, {5}, {2}}
	sort.Slice(s, func(i, j int) bool {
		if s[i].v < s[j].v {
			return false
		}
		return true
	})
	// A
	fmt.Printf("%#v", s)
}
