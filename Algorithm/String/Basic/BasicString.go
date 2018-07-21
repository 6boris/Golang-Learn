package main

import "fmt"

func main() {
	s := "KyleLiu刘家"
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}

	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
}
