package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // memiliki tipe Vertex
	q = &Vertex{1, 2} // memiliki tipe *Vertex
	r = Vertex{X: 1}  // Y:0 di set secara implisit
	s = Vertex{}      // X:0 dan Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}
