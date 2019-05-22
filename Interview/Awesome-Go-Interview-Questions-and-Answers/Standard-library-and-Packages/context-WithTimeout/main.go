package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeout := 10 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	//defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("waited for 1 sec")
	case <-time.After(2 * time.Second):
		fmt.Println("waited for 2 sec")
	case <-time.After(3 * time.Second):
		fmt.Println("waited for 3 sec")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
