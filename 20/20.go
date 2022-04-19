/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

package twenty

import (
	"strings"
)

func Reverse(str string) string {
	words := strings.Fields(str)
	if len(words) <= 1 {
		return str
	}
	res := make([]string, len(words))
	for i, j := len(words)-1, 0; i >= 0; i, j = i-1, j+1 {
		res[i] = words[j]
	}
	resString := strings.Join(res, " ")
	return resString
}
