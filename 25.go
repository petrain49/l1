package main

import "time"

/*
Реализовать собственную функцию sleep.
*/
func mySleepTimer(s int) string {
	t := time.NewTimer(time.Duration(s) * time.Second)
	res := <-t.C
	return res.String()
}


func mySleepAfter(s int) string {
	res := <-time.After(time.Duration(s) * time.Second)
	return res.String()
}

func mySleepSubstraction(s int) string {
	start := time.Now()
	sDur := time.Duration(s)

	for {
		end := time.Now()
		sub := end.Sub(start) / time.Second
		if sub >= sDur {
			return sub.String()
		}
	}
}