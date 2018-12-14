package main

import (
	"crypto/rand"
	"math/big"
	"fmt"
)

func randomTest() {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(1000))
	fmt.Println(randomNumber)

	for i := 0; i < 20; i++ {
		randomNumber, _ := rand.Int(rand.Reader, big.NewInt(20000))
		fmt.Println(randomNumber)
	}
}

func main() {
	randomTest()
}
