package main

import "fmt"

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z.area())
	fmt.Println(z)
}

func main() {
	s := Square{2}
	info(s)
}
