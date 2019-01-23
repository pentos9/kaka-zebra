package main

import "fmt"

type VertexMap struct {
	Lat, Long float64
}

var mp = map[string]VertexMap{
	"Bell Labs": VertexMap{
		40.68433, -74.39967,
	},
	"Google": VertexMap{
		37.42202, -122.08408,
	},
}

func main() {
	mp = make(map[string]VertexMap)
	mp["Bell Labs"] = VertexMap{
		40.68433, -74.39967,
	}
	fmt.Println(mp["Bell Labs"])
}
