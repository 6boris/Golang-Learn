package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"sync"
	"time"
)

func bigBytes() *[]byte {
	s := make([]byte, 100000000)
	return &s
}

func main() {
	var wg sync.WaitGroup

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		for {
			time.Sleep(time.Second * 300)
			debug.FreeOSMemory()
		}
	}()
	for i := 0; i < 1000000; i++ {
		s := bigBytes()
		if s == nil {
			log.Println("oh noes")
		}

	}

	log.Println("go to wait")
	wg.Add(1)
	wg.Wait()
}
