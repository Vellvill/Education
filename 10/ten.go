package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.0, -27.0, -21.0, 13.0, 19.0, 15.5, 24.5, 6.0, -7.0, 4.0}

	groups := make(map[float64][]float64)

	var g float64
	for _, t := range temperatures {
		g = math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}

	for g, t := range groups {
		fmt.Printf("%v: %v\n", g, t)
	}
}
