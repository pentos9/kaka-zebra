package main

import "fmt"

type MyVertex struct {
	Lat, Long float64
}

var m map[string]MyVertex

func main() {
	m = make(map[string]MyVertex)
	m["Bell Labs"] = MyVertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
