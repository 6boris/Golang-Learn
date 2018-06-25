package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(2*i - 1)
		}(i)
	}

	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println(2 * i)
		}(i)
	}

	time.Sleep(3 * time.Second)
}
