package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(len("你好"), utf8.RuneCountInString("你好"))
}
