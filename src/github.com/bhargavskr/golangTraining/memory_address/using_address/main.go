package main

import "fmt"

const metersToyards float64 = 1.09361

func main() {

	var meters float64
	fmt.Print("Enter meter swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToyards
	fmt.Println(meters, " meters is ", yards, "yards.")

}
