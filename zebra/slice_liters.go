package main

import "fmt"

func sliceLiters() {
	q := []int{1, 2, 4, 5, 6, 87}
	fmt.Println(q)

	r := []bool{true, false, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{3, true},
		{56, false},
		{4, true},
		{9, true},
	}

	fmt.Println(s)
}

func bounds() {
	s := []int{5, 6, 7, 2, 12}
	fmt.Println(s)

	s = s[0:3]
	fmt.Println(s)

	s = s[:4]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func main() {
	sliceLiters()
	bounds()
}
