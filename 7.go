package main

import (
	"sync"
)

type sMap struct {
	sync.Mutex
	m map[int]int
}

func (sm *sMap) write(wg *sync.WaitGroup, i int, v int) {
	sm.Lock()
	sm.m[i] = v
	sm.Unlock()
	wg.Done()
}

/*
Реализовать конкурентную запись данных в map.
*/
func seven() map[int]int {
	var wg sync.WaitGroup
	mp := make(map[int]int)
	sm := sMap{m: mp}

	for x := 0; x < 1000000; x++ {
		wg.Add(1)
		go sm.write(&wg, x, x)
	}

	wg.Wait()

	return sm.m
}
