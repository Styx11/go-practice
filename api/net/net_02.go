package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

func check(err error) {
	if err != nil {
		panic("Error at client")
	}
}

// net_02 creates a client to send a message from terminal to server localhost:8080
func main() {
	flag.Parse()
	args := flag.Args()
	msg := strings.Join(args, " ")
	if len(args) <= 0 {
		fmt.Println("Please enter a message you want to send")
		return
	}

	// just dial to :8080
	conn, err := net.Dial("tcp", "localhost:8080")
	check(err)
	defer conn.Close()

	_, err = conn.Write([]byte(msg))
	check(err)

	fmt.Printf("Message have been sent: %s\n", msg)
}
