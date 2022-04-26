/*
Реализовать конкурентную запись данных в map.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type mapa struct {
	m map[int]int
	sync.RWMutex
}

func create() *mapa {
	return &mapa{make(map[int]int), sync.RWMutex{}}
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
	for i := 0; i <= 100000; i++ {
		s := rand.Intn(1000)
		go func(s int) {
			m.RLock()
			fmt.Println(m.m[s])
			m.RUnlock()
		}(s)
	}
}
