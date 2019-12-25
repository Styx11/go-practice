package main

import (
	"fmt"
	"io"
	"net"
)

func check(err error) {
	if err != nil {
		fmt.Println("Error at server:", err.Error())
	}
}

// Writer or bufio.Writer will get stuck
// so we use goruntine
func handleEcho(conn net.Conn) chan []byte {
	ch := make(chan []byte)

	go func() {
		for msg := range ch {
			_, err := conn.Write(msg)
			check(err)
		}
	}()

	return ch
}

// handle connect: receive message from client and start a goruntine to echo back
func handleConn(conn net.Conn) {
	defer conn.Close()

	msgChan := handleEcho(conn)

	fmt.Print("Echo:")
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("")
			close(msgChan)
			break
		}

		fmt.Printf(" %s", buf[:n])
		msgChan <- buf[:n]
	}
}

// an Echo server
func main() {
	ln, err := net.Listen("tcp", ":8083")
	check(err)
	defer ln.Close()

	fmt.Println("Server running at localhost:8083...")
	for {
		conn, err := ln.Accept()
		check(err)

		go handleConn(conn)
	}
}
