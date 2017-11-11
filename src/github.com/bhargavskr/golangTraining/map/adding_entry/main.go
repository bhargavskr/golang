package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good Morning",
		"Jenny": "Bonjour",
	}

	myGreeting["Kris"] = "Me"
	fmt.Println(myGreeting)

}
