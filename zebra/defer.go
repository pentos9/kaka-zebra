package main

import "fmt"

func deferTest ()  {
	defer fmt.Println("It's defer!")
	fmt.Println("[NotDefer] None defer block!")
}

func main() {
	deferTest()
	fmt.Println("hello go2!")
}
