package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)
	//m := sync.Map{}

	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			defer mu.Unlock()

			mu.Lock()
			m[rand.Int()] = rand.Int()
			//m.Store(rand.Int(), rand.Int())
		}()
	}
	wg.Wait()
	fmt.Println(m)
	//m.Range(func(key, value interface{}) bool {
	//	fmt.Println(key, value)
	//	return true
	//})
}
