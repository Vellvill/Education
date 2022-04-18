package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	first()  //using special empty struct channel
	second() //using context.WithTimeOut
	third()
}

func first() {
	ch := make(chan struct{})
	go routine(ch, context.TODO())
	time.Sleep(10 * time.Second)
	ch <- struct{}{}
}

func second() {
	ch := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go routine(ch, ctx)
	time.Sleep(15 * time.Second)
}

func third() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-time.After(10 * time.Second):
			cancel()
		}
	}()
	go routine(ch, ctx)
	time.Sleep(15 * time.Second)
}

func routine(done <-chan struct{}, ctx context.Context) {
	for {
		select {
		case <-done:
			fmt.Println("done channel")
			return
		case <-ctx.Done():
			fmt.Println("done context")
			return
		case <-time.After(2 * time.Second):
			fmt.Println("working")
		}
	}
}
