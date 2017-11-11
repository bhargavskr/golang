package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {

	var p1 Person
	//	p1 := Person{"James", "Bond", 22}
	//	bs, v := json.Marshal(p1)
	fmt.Println(p1)
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":22}`)
	json.Unmarshal(bs, &p1)
	fmt.Printf("%T \n", p1)
	fmt.Println(p1)

}
