/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//inc + Mutex + Waitgroup
type inc struct {
	num int
	sync.Mutex
	sync.WaitGroup
}

//inc2 + Channels
type inc2 struct {
	num         int
	chMutex     chan struct{}
	chWaitGroup chan struct{}
}

//inc3 + Atomic
type inc3 struct {
	num int64
	sync.WaitGroup
}

func newInc() *inc {
	return &inc{0, sync.Mutex{}, sync.WaitGroup{}}
}

func (i *inc) toNil() {
	i.num = 0
}

func (i *inc) increase(n int) {
	i.Add(n)
	for j := 0; j < n; j++ {
		go func() {
			defer i.Done()
			i.Lock()
			i.num++
			i.Unlock()
		}()
	}
	i.Wait()
}

func newInc2() *inc2 {
	return &inc2{0, make(chan struct{}, 1), make(chan struct{})}
}

func (i *inc2) increase(n int) {
	for j := 0; j < n; j++ {
		go func() {
			i.chMutex <- struct{}{}
			i.num++
			<-i.chMutex
			i.chWaitGroup <- struct{}{}
		}()
	}
	for j := 0; j < n; j++ {
		<-i.chWaitGroup
	}
}

func (i *inc2) toNil() {
	i.num = 0
}

func newInc3() *inc3 {
	return &inc3{0, sync.WaitGroup{}}
}

func (i *inc3) increase(n int) {
	i.Add(n)
	for j := 0; j < n; j++ {
		go func() {
			defer i.Done()
			atomic.AddInt64(&i.num, 1)
		}()
	}
	i.Wait()
}

func main() {
	inc := newInc()   //using mutex
	inc2 := newInc2() //using channels
	inc3 := newInc3() //using atomic
	inc.increase(36573)
	inc2.increase(643745)
	inc3.increase(54235)
	fmt.Printf("first: %d\nsecond: %d\nthird: %d\n", inc.num, inc2.num, inc3.num)
}
