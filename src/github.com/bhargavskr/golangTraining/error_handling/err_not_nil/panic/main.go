package main

import (
	"os"
)

func main() {
	_, err := os.Open("me.txt")
	if err != nil {
		panic(err)

	}
}
