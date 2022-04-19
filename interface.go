package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	sideLengh float32
}

func (s Square) Area() float32 {
	return s.sideLengh * s.sideLengh
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return float32(math.Pi) * c.radius
}

func main() {
	square := Square{sideLengh: 10.5}
	circle := Circle{radius: 5.7}

	printShape(square)
	printShape(circle)
}

func printShape(s Shape) {
	fmt.Printf("Площадь фигуры %.2f см\n", s.Area())
}
