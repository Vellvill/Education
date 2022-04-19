/*
Реализовать пересечение двух неупорядоченных множеств.
*/

package main

import "fmt"

func main() {
	array1 := []int{6, 4, 3, 6, 1, 2, 3, 7, 5, 6}
	array2 := []int{6, 8, 4, 2, 1, 6}
	fmt.Println(first(array1, array2))
	fmt.Println(second(array1, array2))
}

//O(n*2)
func first(in1, in2 []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for _, v := range in1 {
		m[v] = 1
	}
	for _, v := range in2 {
		if j, ok := m[v]; ok {
			if j == 1 {
				res = append(res, v)
				m[v]++
			}
		}
	}
	return res
}

//O(n^2)
func second(in1, in2 []int) []int {
	res := make([]int, 0)
	mapa := make(map[int]int)
	for _, up := range in1 {
		for _, down := range in2 {
			if up == down {
				if _, ok := mapa[up]; !ok {
					res = append(res, up)
					mapa[up] = down
				}
			}
		}
	}
	return res
}
