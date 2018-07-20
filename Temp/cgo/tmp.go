package main

/*
   int add(int x, int y) {
       return x + y;
   }
*/
import "C"
import "fmt"

func main() {
	var x C.int = C.add(1, 2)
	var y int = int(x)
	fmt.Println(x, y)
}
