package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height, base float64
}

func main() {
	sq := square{sideLength: 10.0}
	printArea(sq)
	tr := triangle{
		height: 10.0,
		base:   20.0,
	}
	printArea(tr)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
