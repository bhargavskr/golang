package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(" error ", err)

	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("me.txt")
	if err != nil {
		log.Println(" error ", err)

	}
}
