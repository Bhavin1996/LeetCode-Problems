package main

import "fmt"

type shape interface {
	Area() int
}

type rectangle struct {
	length int
	height int
}

func (r rectangle) Area() int {
	return r.height * r.length
}

func Calculate(s shape) {
	a := s.Area()
	fmt.Println(a)
}

func main() {
	rect := rectangle{
		length: 5,
		height: 10,
	}
	Calculate(rect)
}
