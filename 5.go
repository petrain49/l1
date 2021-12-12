package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/
func five() {
	n, err := inputNum()
	if err != nil {
		return
	}

	wnr(n)
}

func wnr(n int) {
	timeOut := time.After(time.Duration(n) * time.Second)
	ch := make(chan int)
	done := false

	go func() {
		for {
			select{
			case <-timeOut:
				done = true
				fmt.Println("Timeout")
				break
			default:
				fmt.Println(<-ch)
			}
		}
	}()

	for x := 1 ; !done; x++ {
		ch <- x
	}
}