/*
Реализовать конкурентную запись данных в map.
*/

package main

import (
	"sync"
)

type mapa struct {
	m map[int]int
	sync.Mutex
}

func create() *mapa {
	return &mapa{make(map[int]int), sync.Mutex{}}
}

func main() {
	m := create()
	for i := 0; i <= 1000; i++ {
		go func(i int) {
			m.Lock()
			m.m[i] = i
			m.Unlock()
		}(i)
	}
}
