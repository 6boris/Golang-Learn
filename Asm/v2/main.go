package main

import "fmt"

func myFunction(i int, arr [2]int) {
	fmt.Printf("in my_funciton - i=%p arr=%p\n", &i, &arr)
}

func main() {
	i := 30
	arr := [2]int{66, 77}
	fmt.Printf("before calling - i=%p arr=%p\n", &i, &arr)
	myFunction(i, arr)
	fmt.Printf("after  calling - i=%p arr=%p\n", &i, &arr)
}
