package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
func four() {
    n, err := inputNum()
    if err != nil {
        return
    }

    work(n)
}

func work(n int) {
    var wg sync.WaitGroup
    jobs := make(chan int, 100)
    done := false

    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        fmt.Fprintln(os.Stdout, "END OF WORK")
        done = true
    }()

    wg.Add(n)
    for x := 1; x <= n; x++ {
        go worker(&wg, x, jobs)
		fmt.Println()
    }

    for j := 1; !done; j++ {
        jobs <- j
    }
    close(jobs)

    wg.Wait()
}

func worker(wg *sync.WaitGroup, id int, jobs <-chan int) {
    for j := range jobs {
        fmt.Fprintf(os.Stdout, "worker №%d did the job №%d\n", id, j)
    }

    fmt.Fprintf(os.Stdout, "worker %d is dead\n", id)
    wg.Done()
}

func inputNum() (int, error) {
    var n int
    _, err := fmt.Scanf("%d", &n)

    return n, err
}