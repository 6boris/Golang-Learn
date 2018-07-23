package main

import "fmt"

func main() {
	num := 3
	fmt.Println(num&(num-1) == 0)

}
