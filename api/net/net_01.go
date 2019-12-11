package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
)

var (
	clients []string       //store client names
	msgReg  *regexp.Regexp //match message from clients exclude client name and 'says: '
	nameReg *regexp.Regexp //match clients
)

func init() {
	msgReg = regexp.MustCompile(`says:\s([!-~\s]*)`) //all inputable characters
	nameReg = regexp.MustCompile(`^([A-Z]+(?:\s[A-Z]+)?)\ssays\:\s`)
}

func check(err error) {
	if err != nil {
		fmt.Println("Error in server:", err.Error())
	}
}

// in ADMIN mode:
// * The client can shut down the server by sending a command SH
// * When the client sends the WHO command, the server displays a list of user names
func handleConn(conn net.Conn) {
	defer conn.Close()

	buf, err := ioutil.ReadAll(conn)
	check(err)
	rawMsg := string(buf)
	clientMsg := msgReg.FindStringSubmatch(rawMsg)[1]
	clientName := nameReg.FindStringSubmatch(rawMsg)[1]

	switch clientMsg {
	case "SH":
		if clientName != "ADMIN" {
			fmt.Println("Only administrator has permission!")
			return
		}

		fmt.Println("SH: We gonna shot it down")
		conn.Close() //defer won't get called when using os.Exit
		os.Exit(0)
	case "WHO":
		if clientName != "ADMIN" {
			fmt.Println("Only administrator has permission!")
			return
		}

		fmt.Println("WHO: This is the client list: 1:active, 0=inactive")
		if len(clients) == 0 {
			fmt.Println("No user yet")
			return
		}
		for _, cli := range clients {
			fmt.Printf("User %s is 1\n", cli)
		}
	default:
		if clientName != "ADMIN" {
			clients = append(clients, clientName)
			fmt.Printf("Received data: %s\n", rawMsg)
		} else {
			fmt.Println("Unsupported admin command:", clientMsg)
		}
	}
}

// this example starts a tcp server receiving message from clients
// rawMsg includes like: JOHN says: hello
// more detail: https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.1.md
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
