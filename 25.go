package main

import "time"

/*
Реализовать собственную функцию sleep.
*/
func mySleep(n int) {
	<-time.After(time.Duration(n) * time.Second)
}