package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var sec int
	fmt.Printf("insert time to work:\n")
	_, _ = fmt.Scan(&sec)
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(sec))
	defer cancel()
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go writer(ctx, ch)
	go reader(ctx, wg, ch)
	wg.Wait()
}

func writer(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time expired, writer")
			return
		default:
			ch <- rand.Intn(100)
		}
	}
}

func reader(ctx context.Context, wg *sync.WaitGroup, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time expired, reader")
			wg.Done()
		case x := <-ch:
			fmt.Println(x)

		}
	}
}
