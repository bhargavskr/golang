package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Fatalln(err)
		}

		io.WriteString(conn, "\n Helllo")
		fmt.Fprintln(conn, "How is yor day?")
		fmt.Fprintln(conn, "%v", "Great")
		conn.Close()
	}
}
