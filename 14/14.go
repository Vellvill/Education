/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	interfaceSlice := make([]interface{}, 0)

	interfaceSlice = append(interfaceSlice, true)
	interfaceSlice = append(interfaceSlice, 50)
	interfaceSlice = append(interfaceSlice, make(chan interface{}))
	interfaceSlice = append(interfaceSlice, "hello")

	defen1(interfaceSlice)
	defen2(interfaceSlice)
}

func defen2(in []interface{}) {
	for i, v := range in {
		elemType := reflect.ValueOf(v).Kind()
		switch elemType {
		case reflect.String:
			fmt.Printf("%d elem of slice is string\n", i)
		case reflect.Int:
			fmt.Printf("%d elem of slice is int\n", i)
		case reflect.Chan:
			fmt.Printf("%d elem of slice is chan\n", i)
		case reflect.Bool:
			fmt.Printf("%d elem of slice is bool\n", i)
		}
	}
}

func defen1(in []interface{}) {
	for _, v := range in {
		switch v.(type) {
		case string:
			fmt.Printf("type %T\n", v)
		case bool:
			fmt.Printf("type %T\n", v)
		case int:
			fmt.Printf("type %T\n", v)
		case chan interface{}:
			fmt.Printf("type %T\n", v)
		}
	}
}
