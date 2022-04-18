package main

import (
	"fmt"
	"sync"
)

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	solutionOne(array)
	solutionTwo(array)
}

func solutionOne(array [5]int) {
	ch := make(chan int)

	for _, v := range array {
		go func(v int) {
			ch <- v * v
		}(v)
	}

	for i := 0; i < len(array); i++ {
		fmt.Printf("first %d\n", <-ch)
	}
}

func solutionTwo(array [5]int) {
	wg := new(sync.WaitGroup)

	wg.Add(len(array))
	for _, v := range array {
		go func(v int) {
			defer wg.Done()
			fmt.Printf("second %d\n", v*v)
		}(v)
	}
	wg.Wait()
}
