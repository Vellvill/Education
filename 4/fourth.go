/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.

*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Intn(100)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	wg := new(sync.WaitGroup)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		select {
		case <-signalChan:
			cancel()
		}
	}()
	var workers int
	fmt.Println("Insert number of workers:")
	_, _ = fmt.Scan(&workers)
	wg.Add(workers)
	for n := 1; n <= workers; n++ {
		go worker(ctx, ch, wg)
	}
	wg.Wait()
}

func worker(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ch:
			fmt.Printf("%d\n", <-ch)
		case <-ctx.Done():
			fmt.Printf("catched ctrl+c signal\n")
			wg.Done()
		}
	}
}
