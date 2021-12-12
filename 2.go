package main

import (
	"fmt"
	"os"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
func two() {
	arr := [...]int{2, 4, 6, 8, 10}

	curRes := make(chan int)
	for _, n := range arr {
		go pow2(n, curRes)
		fmt.Fprintln(os.Stdout, <-curRes)
	}
}

func pow2(n int, res chan int) {
	res <- n * n
}