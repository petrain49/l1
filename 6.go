package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/
func six() {
	var wg sync.WaitGroup

	quit := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(4)
	go firstRoutine(&wg, quit)
	go secondRoutine(&wg, quit)
	go thirdRoutine(&wg, ctx)
	go forthRoutine(&wg)

	quit <- 10
	quit <- 20
	quit <- 30
	close(quit)
	cancel()

	wg.Wait()
}

func firstRoutine(wg *sync.WaitGroup, quit chan int) {
	for {
		select {
		case <-quit:
			fmt.Println("firstRoutine done")
			wg.Done()
			return
		default:
			fmt.Println("firstRoutine working")
		}
	}
}

func secondRoutine(wg *sync.WaitGroup, ch chan int) {
	for {
		_, ok := <-ch
		if !ok {
			fmt.Println("secondRoutine done")
			wg.Done()
			return
		}
		fmt.Println("secondRoutine working")
	}
}

func thirdRoutine(wg *sync.WaitGroup, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("thirdRoutine done")
			wg.Done()
			return
		default:
			fmt.Println("thirdRoutine working")
		}
	}
}

func forthRoutine(wg *sync.WaitGroup) {
	timeOut := time.After(1 * time.Microsecond)
	for {
		select {
		case <-timeOut:
			fmt.Println("forthRoutine done")
			wg.Done()
			return
		default:
			fmt.Println("forthRoutine working")
		}
	}
}