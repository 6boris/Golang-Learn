package main

import (
	"fmt"
	"time"
)


func main() {
	t1 := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2017, time.February, 16, 0, 0, 0, 0, time.UTC)

	fmt.Println(t2.Sub(t1).String())
}
