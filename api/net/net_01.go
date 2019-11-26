package main

import (
	"fmt"
	"io"
	"net"
)

func check(err error) {
	if err != nil {
		fmt.Println("Error in server:", err.Error())
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Received from client:")
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("")
			break
		}
		check(err)

		fmt.Printf(" %s", buf[:n])
	}
}

// this example starts a tcp server receiving message from clients
func main() {
	ln, err := net.Listen("tcp", ":8080")
	check(err)
	defer ln.Close()

	fmt.Println("Start server at localhost:8080...")
	for {
		conn, err := ln.Accept()
		check(err)

		go handleConn(conn)
	}
}
