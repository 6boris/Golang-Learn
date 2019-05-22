package main

import "fmt"

type S struct {
	name string
}

func main() {
	m := map[string]*S{"x": &S{"one"}}
	//m := map[string]S{"x": S{"one"}}
	m["x"].name = "two"
	fmt.Println(m["x"].name)
}
