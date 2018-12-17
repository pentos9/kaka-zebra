package main

import (
	"fmt"
	"math"
)

func infiniteLoop() {
	for {
		fmt.Println("infiniteLoop")
	}
}

func whileLoop() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func print(text string) {
	fmt.Println(text)
}

func main() {
	whileLoop()
	print(sqrt(6.67))
	fmt.Println(sqrt(2), sqrt(-4))
}
