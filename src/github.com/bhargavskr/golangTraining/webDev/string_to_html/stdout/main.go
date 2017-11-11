package main

import "fmt"

func main() {

	name := "Bhargav"

	tpl := `
		<html>
		<head><title>New Page</title></head>
		<body><h2>` + name + `</h2></body>
		</html>
	`
	fmt.Println(tpl)

}
