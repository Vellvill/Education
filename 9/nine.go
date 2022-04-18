package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		defer close(ch1)
		for _, v := range array {
			ch1 <- v
		}
	}()

	go func() {
		defer close(ch2)
		for v := range ch1 {
			ch2 <- v * 2
		}
	}()

	for v := range ch2 {
		fmt.Println(v)
	}
}
