package main

import (
	"fmt"
)

func SwitchOnType(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("INT")
	case string:
		fmt.Println("String")
		//case Contact:
	default:
		fmt.Println("No Type Match")

	}
}

func main() {
	SwitchOnType(8)
	SwitchOnType("Me")

}
