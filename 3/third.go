package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	first()
	second()
	third()
}

func first() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer close(ch1)
		for i := 2; i <= 10; i += 2 {
			ch1 <- i
		}
	}()
	go func() {
		defer close(ch2)
		for i := range ch1 {
			ch2 <- i * i
		}
	}()
	var num int
	for v := range ch2 {
		num += v
	}
	fmt.Printf("first:%d\n", num)
}

func second() {
	array := []int32{2, 4, 6, 8, 10}
	var num int32
	wg := new(sync.WaitGroup)
	wg.Add(len(array))
	for _, v := range array {
		go func(v int32) {
			defer wg.Done()
			atomic.AddInt32(&num, v*v)
		}(v)
	}
	wg.Wait()
	fmt.Printf("second:%d\n", num)
}

func third() {
	array := []int{2, 4, 6, 8, 10}
	mu := make(chan struct{}, 1)
	wg := make(chan struct{})
	var num int
	increase := func(v int) {
		mu <- struct{}{}
		num += v * v
		<-mu
	}
	for _, v := range array {
		go func(v int) {
			increase(v)
			wg <- struct{}{}
		}(v)
	}
	for i := 0; i < 5; i++ {
		<-wg
	}
	fmt.Printf("third:%d\n", num)
}
