package main

import (
	"fmt"
	"time"

	"github.com/EDDYCJY/gsema"
)

var sema = gsema.NewSemaphore(3)

func main() {
	userCount := 10
	for i := 0; i < userCount; i++ {
		go Read(i)
	}

	sema.Wait()
}

func Read(i int) {
	defer sema.Done()
	sema.Add(1)

	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
}
