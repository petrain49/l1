package main

import (
	"fmt"
	"os"
	"sync"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/
func pipeline() {
	var wg sync.WaitGroup

	fch := make(chan int)
	sch := make(chan int)

	wg.Add(2)
	go f(&wg, fch)
	go s(&wg, fch, sch)

	for {
		x, ok := <-sch
		if !ok {
			break
		}
		fmt.Fprintln(os.Stdout, x)
	}

	wg.Wait()
}

func f(wg *sync.WaitGroup, ch chan<- int) {
	for x := 0; x < 100; x++ {
		ch <- x
	}
	close(ch)
	wg.Done()
}

func s(wg *sync.WaitGroup, input <-chan int, output chan<- int) {
	for {
		x, ok := <-input
		if !ok {
			break
		}
		output <- x * 2
	}
	close(output)
	wg.Done()
}
