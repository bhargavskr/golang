package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("me.txt")
	if err != nil {
		log.Println(" error ", err)

	}
}
