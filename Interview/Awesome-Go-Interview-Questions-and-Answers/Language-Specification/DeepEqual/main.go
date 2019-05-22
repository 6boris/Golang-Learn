package main

import (
	"fmt"
	"reflect"
)

type S struct {
	a, b, c string
}

func main() {
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})

	fmt.Println(&x, &y)
	fmt.Println(reflect.DeepEqual(x, y))

	fmt.Println(x == y) //A
}
