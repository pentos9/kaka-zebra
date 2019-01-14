package main

import "fmt"

func printfSlice(s string, x []int) {
	fmt.Printf("%s len=%d,cap=%d %v \n", s, len(x), cap(x), x)
}

func main() {
	a := make([]int, 5)
	fmt.Println(a)

	printfSlice("a", a)

	b := make([]int, 0, 5)
	printfSlice("b", b)

	c := b[2:5]
	printfSlice("c", c)
}
