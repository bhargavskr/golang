package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrMsg = errors.New("square root of negative number")

func main() {
	fmt.Printf("%T \n", ErrMsg)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMsg
	}
	return 42, nil
}
