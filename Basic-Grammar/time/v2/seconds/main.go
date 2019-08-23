package main

import (
	"fmt"
	"time"
)

func main() {
	m, _ := time.ParseDuration("1h1m30s")
	fmt.Printf("take off in t-%.0f seconds.", m.Seconds())
}
