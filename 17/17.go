/*
Бинарный поиск.
*/

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	is, pos := search(array, 3)
	fmt.Println(is, pos)
}

func search(a []int, t int) (bool, int) {
	first, last := 0, len(a)-1
	for first <= last {
		mediana := (first + last) / 2
		if a[mediana] == t {
			return true, a[mediana]
		} else if t > a[mediana] {
			first = mediana + 1
		} else {
			last = mediana - 1
		}
	}
	return false, -1
}
