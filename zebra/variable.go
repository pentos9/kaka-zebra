package main

import "fmt"
import (
	"math/rand"
	"time"
)

func variable() {
	var a = "var"
	var b string = "f-string"
	var c bool
	fmt.Println(a, b, c)

	v1, v2, v3 := "zebra", 231, 4546
	fmt.Println(v1, v2, v3)
}

func random() {
	fmt.Println("random test")
	var random = rand.Intn(1000)
	fmt.Println(random)
	rand.NewSource(time.Now().Unix())
	fmt.Println(rand.Intn(1000))
	fmt.Println(rand.Intn(1000))
	fmt.Println(rand.Intn(1000))
	fmt.Println(rand.Intn(1000))

	fmt.Println(rand.Int())
}

func run() {
	var number int = 12
	fmt.Println(number)

	number = 1500
	fmt.Println(number)

	newNumber := 567
	fmt.Println(newNumber)

}

func main() {
	run()
	variable()
	random()
}
