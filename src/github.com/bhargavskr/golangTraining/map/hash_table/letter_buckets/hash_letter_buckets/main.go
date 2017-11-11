package main

import (
	"bufio"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		log.Fatal(err)
	}

	buckets := make([]int, 200)

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}

	for i, a := range buckets {
		println(string(i), " : ", a)

	}

}

func hashBucket(word string) int {

	return int(word[0])
}
