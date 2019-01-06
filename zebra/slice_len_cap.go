package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d,%v \n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[0:3]
	printSlice(s)

	s = s[:5]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	var array []int

	fmt.Println(s, len(array), cap(array))
	if array == nil {
		fmt.Println("nil!")
	}
}
