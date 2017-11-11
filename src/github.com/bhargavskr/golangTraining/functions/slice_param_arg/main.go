package main

import "fmt"

func average(sf []float64) float64 {

	fmt.Println(sf)
	fmt.Printf("%T \n", sf)

	total := 0.0
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

func main() {
	data := []float64{12, 23, 44, 3343, 2, 335, 44}
	fmt.Println(average(data))
}
