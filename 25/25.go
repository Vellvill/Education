/*
Реализовать собственную функцию sleep
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var t time.Time

func init() {
	t = time.Now()
}

func main() {
	var seconds int
	fmt.Println("введите время сна основного потока в секундах: ")
	_, _ = fmt.Scan(&seconds)
	firstSolution(seconds)
	secondSolution(seconds)
}

func firstSolution(seconds int) {
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Printf("Второстпенный поток, t: %d\n", int(time.Since(t).Seconds()))
			case <-done:
				return
			}
		}
	}()
	mysleep(seconds)
	fmt.Println("Основной поток проснулся")
	done <- struct{}{}
}

func mysleep(seconds int) {
	fmt.Println("основной поток уснул")
	<-time.After(time.Second * time.Duration(seconds))
}

func secondSolution(seconds int) {
	start := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(seconds)
	go func() {
		for {
			select {
			case <-start:
				sleep(seconds, wg)
			}
		}
	}()
	fmt.Println("Основной поток уснул")
	start <- struct{}{}
	wg.Wait()
	fmt.Println("Основной поток проснулся")
}

func sleep(n int, wg *sync.WaitGroup) {
	for i := n; i >= 0; i-- {
		fmt.Printf("Второстепенный поток, t: %d\n", int(time.Since(t).Seconds()))
		<-time.After(1 * time.Second)
		wg.Done()
	}
}
