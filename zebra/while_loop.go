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

func squareRoot(x float64) string {
	if x < 0 {
		return squareRoot(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func print(text string) {
	fmt.Println(text)
}

func main() {
	whileLoop()
	print(squareRoot(6.67))
	fmt.Println(squareRoot(2), squareRoot(-4))
}
