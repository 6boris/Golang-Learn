package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Kyel刘家"

	for i, ch := range s {
		fmt.Printf("%d %X\n", i, ch)
	}
	fmt.Println(utf8.RuneCountInString(s))
}
