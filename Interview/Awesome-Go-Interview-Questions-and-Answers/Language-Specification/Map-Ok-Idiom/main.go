package main

import "fmt"

func main() {
	var m map[string]int //A
	//m := make(map[string]int)

	m["a"] = 1

	if v := m["b"]; v != nil { //B
		fmt.Println(v)
	}

	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}
}
