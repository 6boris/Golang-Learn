package Singleton

import "testing"

// 单例模式懒汉模式
func TestSingletonLazy(t *testing.T) {
	ins1 := GetInstanceLazy()
	ins2 := GetInstanceLazy()
	if ins1 != ins2 {
		t.Fatal("Instance is not equal")
	}
}

//	单列模式饿汉模式
func TestSingletonEager(t testing.T) {
	ins1 := GetIns()
	ins2 := GetIns()
	if ins1 != ins2 {
		t.Fatal("Instance is not equal")
	}
}
