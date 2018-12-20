package main

import (
	"fmt"
	"runtime"
)

func switchGo() {
	fmt.Println("Go runs On")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s.\n", os)
	}

}

func main() {
	switchGo()
}
