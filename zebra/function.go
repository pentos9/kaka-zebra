package main

import "fmt"

var c, python, java, golang bool

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(89, 65))
	fmt.Println(swap("namibia", "pou"))
	fmt.Println(split(900))
}
