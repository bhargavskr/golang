package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)

	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		bufio.NewScanner(conn)
		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected")

		conn.Close()
	}

}
