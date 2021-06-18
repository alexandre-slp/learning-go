package main

import (
	"fmt"
	"math"
)

type square struct {
	l float64
}

type circle struct {
	r float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.l * s.l
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	square := square{
		l: 5,
	}

	circle := circle{
		r: 2,
	}

	fmt.Println(square.area())
	fmt.Println(circle.area())
	info(square)
	info(circle)
}
