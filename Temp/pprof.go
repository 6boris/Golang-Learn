package main
import (
	_ "net/http/pprof"
	"net/http"
	"time"
)
func main() {
	go http.ListenAndServe("localhost:6060", nil)
	for {
		time.Sleep(time.Second)
	}
}