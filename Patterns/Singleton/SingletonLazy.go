package Singleton

import "sync"

//Singleton 是单例模式类
type SingletonLazy struct{}

var singleton *SingletonLazy
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstanceLazy() *SingletonLazy {
	once.Do(func() {
		singleton = &SingletonLazy{}
	})

	return singleton
}
