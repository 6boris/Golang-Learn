package main

import "fmt"

func less(a int, b int) bool {
	return a < b
}

func main() {
	fmt.Println(less(1, 2))
}
