package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := createHugeString(1 << 20)
	fmt.Printf("inset cutter, it should be between 0 and %d\n", len(a))
	var cutter int
	_, _ = fmt.Scan(&cutter)
	if cutter >= len(a) {
		fmt.Println("too big")
		return
	} else {
		ns := justString(cutter, a)
		fmt.Printf("new string is %s", ns)
	}
}

func justString(cutter int, a string) string {
	return a[:cutter]
}

func createHugeString(n int) string {
	res := make([]byte, n)
	for n := range res {
		res[n] = byte(rand.Intn(122-48) + 48)
	}
	return string(res)
}
