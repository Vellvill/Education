package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	delete(2, a)
}

func delete(i int, ar []int) {
	ar[i] = ar[len(ar)-1] //копируем последний элемент в индекс i
	ar = ar[:len(ar)-1]   //усекаем срез
	fmt.Println(ar)
}

