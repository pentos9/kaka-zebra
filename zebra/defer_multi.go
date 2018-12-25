package main

import "fmt"

func deferMulti() {
	for i := 0; i < 2000; i++ {
		defer fmt.Println(i)
	}
}

func deferStatement() {
	fmt.Println("no deffer block!")
	defer fmt.Println("It's defer now!")
}

func main() {
	deferMulti()
	deferStatement()
}
