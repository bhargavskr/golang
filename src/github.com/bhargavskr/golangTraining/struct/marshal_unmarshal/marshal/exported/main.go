package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 22, 007}
	bs, v := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(v)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))

}
