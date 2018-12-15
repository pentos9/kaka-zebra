package main

import "fmt"

func swapString(x, y string) (string, string, string) {
	return y, ",", x
}

func main() {
	fmt.Print(swapString("angy", "busato"))
}
