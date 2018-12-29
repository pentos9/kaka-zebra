package main

import "fmt"

type VertexDTO struct {
	X, Y int
}

var (
	v1 = VertexDTO{1, 3}
	v2 = VertexDTO{X: 34}
	v3 = VertexDTO{}
	p  = &VertexDTO{1, 34}
)

func main() {
	fmt.Println(v1, v2, v3, p)
	fmt.Println(*p)
}
