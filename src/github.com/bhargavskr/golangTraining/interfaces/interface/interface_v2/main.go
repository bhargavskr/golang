package main

import "fmt"
import "math"

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z Shape) {
	fmt.Println(z.area())
	fmt.Println(z)
}

func main() {
	s := Square{2}
	c := Circle{5}
	info(c)
	info(s)
}
