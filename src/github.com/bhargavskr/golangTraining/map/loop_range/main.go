package main

import "fmt"

func main() {

	greeting := map[int]string{

		0: "Good Morning",
		1: "Bonjour!",
		2: "dias!",
		3: "Bon dias!",
	}

	for key, val := range greeting {
		fmt.Println(key, " = ", val)

	}

}

// ALan Donovan, William Kennedy, the w ay to Go
