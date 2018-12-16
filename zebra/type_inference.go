package main

import "fmt"

func typeInference() {
	v := 4.6
	fmt.Printf("v is type of %T", v)
}

func main() {
	typeInference()
}
