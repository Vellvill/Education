package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(solution(strings))
}

func solution(arr []string) []string {
	m := make(map[string]int)
	res := make([]string, 0)
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		} else {
			res = append(res, v)
			m[v] = 1
		}
	}
	return res
}
