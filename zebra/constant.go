package main

import (
	"fmt"
	"crypto/rand"
	"math/big"
)

const PI = 3.14

func main() {
	fmt.Println(PI)

	const Bigger = "Bigger and Bigger"
	fmt.Println(Bigger)
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(8000))
	fmt.Println(randomNumber)
}
