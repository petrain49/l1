package main

import (
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type counter struct {
	sync.Mutex
	num int
}

func (c *counter) inc() {
	c.Lock()
	c.num++
	c.Unlock()
}
