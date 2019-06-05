package o1

import "container/list"

type LRU struct {
	data *list.Element
}

type Data struct {
	Key int
	Val int
}

const N = 10

func (l *LRU) LRUCache() {
}

func (l *LRU) Get(key int) int {
	head := l.data
	for head != nil {
		if head.Value == key {
			return head.Value.(int)
		}
		head = head.Next()
	}
	return 0
}

func (l *LRU) Put(key, val int) {
	data := list.Element{Value: Data{Key: key, Val: val}}
	data.Next() = l.data
	l.data = &data
}
