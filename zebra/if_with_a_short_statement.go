package main

import (
	"math"
	"fmt"
)

func power(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(power(1.5, 2, 100))
	fmt.Println(power(100.5, 2, 100))
}
