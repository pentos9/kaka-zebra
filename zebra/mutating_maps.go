package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	var FUNCTION = "FUNCTION"
	m[FUNCTION] = 2112
	fmt.Println(m)

	m[FUNCTION] = 35
	fmt.Println(m)

	elem, ok := m[FUNCTION]
	fmt.Println(elem, ok)

	elem, ok = m["NotExistKey"]
	fmt.Println(elem, ok)
}
