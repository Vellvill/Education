/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

package main

import "fmt"

func main() {
	var in string
	fmt.Println("insert text")
	_, _ = fmt.Scan(&in)
	res := []rune(in)
	for i, j := 0, len([]rune(res))-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	fmt.Printf("reversed: %s\n", string(res))
}
