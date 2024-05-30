package main

import (
	"fmt"
	"sync"
)

type Trip string

type Driver struct {
	trips []Trip
}

type Stats struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewStats() *Stats {
	return &Stats{
		counters: make(map[string]int),
	}
}

// SetTrips
//  @Description: slice做入参
//  @receiver d
//  @param trips
func (d *Driver) SetTrips(trips []Trip) {
	/*
		Bad
		由于d.trips存储了对入参slice的引用
		而slice包含了指向底层数据的指针
		所以对入参trips的修改也会同步影响d.trips的内容
	*/
	//d.trips = trips

	// Good
	d.trips = make([]Trip, len(trips))
	copy(d.trips, trips)
}

// Snapshot
//  @Description: map 做返回值
//  @receiver s
//  @return map[string]int
func (s *Stats) Snapshot() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	/*
		Bad
		snapshot不再受互斥锁保护
		对snapshot的任何访问都将受到数据竞争的影响
		影响 stats.counters
	*/
	return s.counters

	// Good
	//res := make(map[string]int, len(s.counters))
	//for k, v := range s.counters {
	//	res[k] = v
	//}
	//return res
}

func main() {
	// slice 作为入参
	trips := []Trip{"巴塞罗那", "巴黎", "罗马"}
	d1 := Driver{}
	d1.SetTrips(trips)
	fmt.Println("入参trips修改之前:", d1.trips[0])

	trips[0] = "阿姆斯特丹"
	fmt.Println("入参trips修改之后:", d1.trips[0])

	// map 作为返回值
	stats := NewStats()
	snapshot := stats.Snapshot()
	fmt.Println("用户修改返回值map之前:", stats.counters)
	snapshot["柯南"] = 16
	fmt.Println("用户修改返回值map之后:", stats.counters)
}
