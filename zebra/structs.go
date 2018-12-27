package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Student struct {
	id    int
	name  string
	phone string
}

func main() {
	fmt.Println(Vertex{1, 4})
	fmt.Println(Student{456, "Ferdinand", "magellan"})
}
