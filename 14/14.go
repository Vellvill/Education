package main

import (
	"fmt"
)

func main() {
	var i interface{}
	i = "t43234"
	defen(i)
	i = true
	defen(i)
	i = 55
	defen(i)
	i = make(chan interface{})
	defen(i)
}

func defen(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("type %T\n", i)
	case bool:
		fmt.Printf("type %T\n", i)
	case int:
		fmt.Printf("type %T\n", i)
	case chan interface{}:
		fmt.Printf("type %T\n", i)
	}
}
