package main

func main(){
	// lazy-unsafe-test
	for i := 0;i<1000;i++{
		//go GetLazy()		// 非安全模式下，实例会被创建很多次
		go GetLazySafe()	// 安全模式
		//go GetLazyOnce()	// go专属安全模式
	}
}

// 单元测试
//package singleton
//
//import (
//"sync"
//"testing"
//)
//
//const parCount = 100
//
//func TestSingleton(t *testing.T) {
//	ins1 := GetInstance()
//	ins2 := GetInstance()
//	if ins1 != ins2 {
//		t.Fatal("instance is not equal")
//	}
//}
//
//func TestParallelSingleton(t *testing.T) {
//	wg := sync.WaitGroup{}
//	wg.Add(parCount)
//	instances := [parCount]*Singleton{}
//	for i := 0; i < parCount; i++ {
//		go func(index int) {
//			instances[index] = GetInstance()
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	for i := 1; i < parCount; i++ {
//		if instances[i] != instances[i-1] {
//			t.Fatal("instance is not equal")
//		}
//	}
//}
//
//func TestSingleton2(t *testing.T) {
//	ins1 := GetInstance2()
//	ins2 := GetInstance2()
//	if ins1 != ins2 {
//		t.Fatal("instance is not equal")
//	}
//}
//
//func TestParallelSingleton2(t *testing.T) {
//	wg := sync.WaitGroup{}
//	wg.Add(parCount)
//	instances := [parCount]*Singleton2{}
//	for i := 0; i < parCount; i++ {
//		go func(index int) {
//			instances[index] = GetInstance2()
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	for i := 1; i < parCount; i++ {
//		if instances[i] != instances[i-1] {
//			t.Fatal("instance is not equal")
//		}
//	}
//}

