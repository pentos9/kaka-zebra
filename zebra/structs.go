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
	fmt.Println(Student{456, "Ferdinand", "Magellan"})

	student := Student{1232, "Gosling", "Magellan"}
	fmt.Println(student.name)

	vertex := Vertex{12, 567}
	fmt.Println(vertex.X)
}
