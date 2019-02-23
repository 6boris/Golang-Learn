package main

import "fmt"

//	尝试 chan

func main() {
	userCount := 10
	ch := make(chan bool, 2)
	for i := 0; i < userCount; i++ {
		ch <- true
		go Read(ch, i)
	}

	//time.Sleep(time.Second)
}

func Read(ch chan bool, i int) {
	fmt.Printf("go func: %d\n", i)
	<-ch
}
