package main

import (
	"fmt"
	"math"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
}
