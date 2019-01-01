package main

import "fmt"

func slices() {
	primes := [6]int{1, 3, 4, 5, 6, 7,}
	fmt.Println(primes)

	var s []int = primes[1:5]
	fmt.Println(s)

	names := [4]string{
		"O Henry", "Jack Man", "London", "White Fang",
	}

	fmt.Println(names)

	a := names[1:4]
	b := names[0:3]
	fmt.Println(a, b)
}

func main() {
	slices()
}
