package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")

	if err != nil {
		log.Fatal(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(bs))

}
