package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i, 1)
	}

	for i := 0; i < 5; i++ {
		defer func() {
			defer fmt.Println(i, 2)
		}()
	}

	for i := 0; i < 5; i++ {
		defer func() {
			j := i
			fmt.Println(j, 3)
			fmt.Println(j, &j, i, &i, 3)
		}()
	}

	for i := 0; i < 5; i++ {
		j := i
		defer fmt.Println(j, 4)
	}

	for i := 0; i < 5; i++ {
		j := i
		defer func() {
			fmt.Println(j, 5)
		}()
	}

	for i := 0; i < 5; i++ {
		defer func(j int) {
			fmt.Println(j, 6)
		}(i)
	}

}
