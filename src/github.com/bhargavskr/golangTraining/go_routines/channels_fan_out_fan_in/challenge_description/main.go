package main

import (
	"fmt"
	"time"
)

var workerId int
var publisherId int

func main() {

	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1 * time.Millisecond)

}

func workerProcess(in <-chan string) {
	workerId++
	thisID := workerId
	for {

		fmt.Printf(" %d:  is waiting for input ...  \n ", thisID)
		input := <-in
		fmt.Printf(" %d input is: %s  \n ", thisID, input)

	}

}

func publisher(out chan string) {
	publisherId++
	thisID := publisherId
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data \n ", thisID)
		data := fmt.Sprintf("Data from publisher %d, Data %d ", thisID, dataID)
		out <- data
	}
}
