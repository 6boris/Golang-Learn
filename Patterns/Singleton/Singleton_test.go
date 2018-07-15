package Singleton

import "testing"

func TestSingletonBasic(t *testing.T) {

	// 懒汉模式
	t.Run("TesSingletonLazy...", func(t *testing.T) {
		ins1 := GetInstanceSingletonLazy()
		ins2 := GetInstanceSingletonLazy()
		if ins1 != ins2 {
			t.Fatal("instance is not equal")
		}
	})

	// 饿汉模式
	t.Run("SingletonEager...", func(t *testing.T) {
		ins1 := GetInstanceSingletonEager()
		ins2 := GetInstanceSingletonEager()
		if ins1 != ins2 {
			t.Fatal("instance is not equal")
		}
	})

	// 懒汉加锁方式
	t.Run("SingletonEager...", func(t *testing.T) {
		ins1 := GetInstanceSingletonLazyLock()
		ins2 := GetInstanceSingletonLazyLock()
		if ins1 != ins2 {
			t.Fatal("instance is not equal")
		}
	})

	// 双重锁模式
	t.Run("SingletonDoubleLock...", func(t *testing.T) {
		ins1 := GetInstanceSingletonDoubleLock()
		ins2 := GetInstanceSingletonDoubleLock()
		if ins1 != ins2 {
			t.Fatal("instance is not equal")
		}
	})

	// sync.Once
	t.Run("SingletonOnce...", func(t *testing.T) {
		ins1 := GetInstanceSingletonOnce()
		ins2 := GetInstanceSingletonOnce()
		if ins1 != ins2 {
			t.Fatal("instance is not equal")
		}
	})

}
