package main

import "fmt"

func main() {

	checkThis(7, 7)

}

func checkThis(r, s int) int {
	var u int
	fmt.Println(r, s)
	if r <= 0 || s <= 0 {
		u = r + s
	} else if r > s {
		u = checkThis(r-5, s-4) + s
	} else {
		u = checkThis(r-4, s-5) + r
	}
	fmt.Println(u)
	return u
}
