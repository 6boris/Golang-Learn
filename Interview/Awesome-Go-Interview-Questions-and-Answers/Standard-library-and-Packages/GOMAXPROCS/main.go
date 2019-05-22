package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	const GOMAXPROCS = 1 //A
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for char := 0; char < 26; char++ {
			fmt.Printf("%c ", 'A'+char)
		}
	}()
	go func() {
		defer wg.Done()
		for char := 0; char < 26; char++ {
			fmt.Printf("%c ", 'a'+char)
		}
	}()

	wg.Wait()
}
