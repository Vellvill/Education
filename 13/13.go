/*
Поменять местами два числа без создания временной переменной.
*/

package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var x, y int32 = 1, 2
	fmt.Printf("Before x: %d, y: %d\n", x, y)
	x, y = y, x
	fmt.Printf("After x: %d, y: %d\n", x, y)
	y = atomic.SwapInt32(&x, y)
	fmt.Printf("After x: %d, y: %d\n", x, y)
}
