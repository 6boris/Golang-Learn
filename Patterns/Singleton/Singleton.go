package Singleton

import "sync"

/***************************  懒汉方式  ********************************/

type SingletonLazy struct{}

var insSingletonLazy *SingletonLazy

// 获取懒汉模式对象
func GetInstanceSingletonLazy() *SingletonLazy {
	if insSingletonLazy == nil {
		insSingletonLazy = &SingletonLazy{}
	}
	return insSingletonLazy
}

/***************************  饿汉方式  ********************************/

type SingletonEager struct{}

var insSingletonEager *SingletonEager = &SingletonEager{}

// 获取饿汉模式对象
func GetInstanceSingletonEager() *SingletonEager {
	return insSingletonEager
}

/***************************  懒汉加锁方式  ********************************/

type SingletonLazyLock struct{}

var insSingletonLazyLock *SingletonLazyLock
var muSingletonLazyLock sync.Mutex

func GetInstanceSingletonLazyLock() *SingletonLazyLock {
	muSingletonLazyLock.Lock()
	defer muSingletonLazyLock.Unlock()
	if insSingletonLazyLock == nil {
		insSingletonLazyLock = &SingletonLazyLock{}
	}
	return insSingletonLazyLock
}

/***************************  双重锁方式  ********************************/

type SingletonDoubleLock struct{}

var insSingletonDoubleLock *SingletonDoubleLock
var muSingletonDoubleLock sync.Mutex

func GetInstanceSingletonDoubleLock() *SingletonDoubleLock {
	muSingletonDoubleLock.Lock()
	defer muSingletonDoubleLock.Unlock()
	if insSingletonDoubleLock == nil {
		insSingletonDoubleLock = &SingletonDoubleLock{}
	}
	return insSingletonDoubleLock
}

/***************************  sync.Once  ********************************/

type SingletonOnce struct{}

var insSingletonOnce *SingletonOnce
var once sync.Once

func GetInstanceSingletonOnce() *SingletonOnce {
	once.Do(func() {
		insSingletonOnce = &SingletonOnce{}
	})
	return insSingletonOnce
}
