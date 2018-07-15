package Singleton

type SingletonEager struct{}

var ins *SingletonEager = &SingletonEager{}

func GetIns() *SingletonEager {
	return ins
}
