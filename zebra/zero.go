package main

import "fmt"

func zero() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q", i, f, b, s)

}

func main() {
	zero()
}
