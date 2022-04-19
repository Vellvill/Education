package tventy_six

import (
	"fmt"
	"strings"
)

func IsUnique(str string) bool {
	str = strings.ToLower(str)
	mapa := make(map[rune]int)
	for _, v := range []rune(str) {
		if _, ok := mapa[v]; ok {
			fmt.Printf("%s is not unique\n", string(v))
			return false
		}
		mapa[v] = 1
	}
	return true
}
