package main

import (
	"container/list"
	"fmt"
	"sync"
)

func main() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	l2 := sync.RWMutex{}
	l2.Lock()
	l2.Unlock()
}
