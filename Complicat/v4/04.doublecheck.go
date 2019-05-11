package doublecheck

import (
	"sync"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

/**
A: 不能编译
B: 可以编译，正确实现了单例
C: 可以编译，有并发问题，f函数可能会被执行多次
D: 可以编译，但是程序运行会panic
*/
