package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d", "e", "f"}

	deleteFast(1, a)
}

func deleteFast(i int, ar []string) {
	ar[i] = ar[len(ar)-1] //копируем последний элемент в индекс i
	ar = ar[:len(ar)-1]   //усекаем срез
	fmt.Println(ar)
}

func deleteSlow(i int, ar []string) {
	copy(ar[i:], ar[i+1:]) //сдвиг влево на один индекс
	ar = ar[:len(ar)-1]    //усечь
	fmt.Println(ar)
}
