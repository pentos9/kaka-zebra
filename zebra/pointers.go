package main

import "fmt"

func pointers() {
	var p *int
	fmt.Println(p)

	integer := 98
	p = &integer //point to operand,p stores the address of variable p
	fmt.Println(p)

	fmt.Println(*p) // get the real operand
}

func main() {
	pointers()
}
