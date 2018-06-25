package main

import "fmt"

type Animal struct {
	Name string
	Age  uint
}

func (a Animal) Dmo() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	fmt.Println("Hello Word")
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
}
