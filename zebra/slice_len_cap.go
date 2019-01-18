package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d,%v \n", len(s), cap(s), s)
}

func doPower() {
	var power = []int{1, 3, 4, 5, 5, 6, 6,}
	fmt.Println(power)

	for i := range power {
		power[i] = 1 << uint(i) //resign value
	}

	fmt.Println(power)

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

	array = append(array, 98)
	fmt.Println(array)

	doPower()
}
