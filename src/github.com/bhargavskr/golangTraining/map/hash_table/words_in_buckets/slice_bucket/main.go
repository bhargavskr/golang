package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")

	if err != nil {
		log.Fatal(err)
	}

	buckets := make([][]string, 12)

	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word)
		buckets[n] = append(buckets[n], word)
	}

	for i, a := range buckets {
		fmt.Println(i, " : ", len(a))

	}
	//	fmt.Println(buckets[1])
}

func hashBucket(word string) int {
	v := 0
	for _, a := range word {
		v += int(a)
	}
	return v % 12
}
