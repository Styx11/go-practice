package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic("Error at client")
	}
}

// net_02 creates a client to send a message from terminal to server localhost:8080
// first: user name
// then: message you want to send
// more details: https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.1.md
func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Firstly, what's your name? (only upper name): ")
	rawName, err := inputReader.ReadString('\n')
	check(err)
	clientName := strings.ToUpper(strings.Trim(rawName, "\n"))

	// just dial to :8080
	conn, err := net.Dial("tcp", "localhost:8080")
	check(err)
	defer conn.Close()

	var rawMsg []byte
	for {
		fmt.Print("Enter message (press Q to finish): ")
		input, err := inputReader.ReadBytes('\n')
		check(err)

		input = bytes.Trim(input, "\n") //trim '\n'

		if string(input) == "Q" {
			clientMsg := clientName + " says: " + string(rawMsg)
			_, err := conn.Write([]byte(clientMsg))
			check(err)
			fmt.Printf("Message has been sent: %s\n", clientMsg)
			return
		}

		rawMsg = append(rawMsg, input...)
	}
}
