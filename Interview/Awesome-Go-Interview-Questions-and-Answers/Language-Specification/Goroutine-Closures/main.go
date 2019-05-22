package main

import (
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			defer mu.Unlock()

			mu.Lock()
			m[i] = i
		}(i)

	}
	wg.Wait()
	println(len(m))
}
