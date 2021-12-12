package main

import (
)

/*
Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
func setBit(num int64, i uint, one bool) int64 {
	if one {
    	num |= (1 << i)
	} else {
    	num &= ^(1 << i)
	}

	return num
}