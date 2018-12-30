package main

import "fmt"

func main() {
	var array [2]string
	array[0] = "go"
	array[1] = "lang"
	fmt.Println(array)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
