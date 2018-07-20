package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go http.ListenAndServe("localhost:6060", nil)
	for {
		time.Sleep(time.Second)
	}
}
