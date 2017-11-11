package main

import "fmt"
import "os"
import "log"
import "io"
import "strings"

func main() {

	name := "Bhargav"

	tpl := fmt.Sprintf(`
		<html>
		<head><title>New Page</title></head>
		<body><h2>` + name + `</h2></body>
		</html>
	`)

	nf, err := os.Create("index.html")

	if err != nil {
		log.Fatal("Error in creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
	fmt.Println(tpl)

}
