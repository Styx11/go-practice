package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

var (
	msg  []byte
	port string
	reg  *regexp.Regexp
)

func init() {

	// rest of args are the messages you want to send to server so flag isn't good enough
	// match port like: -p :80, --p=:80, --port=:8080 or --port :8080
	var rawMsg string
	args := strings.Join(os.Args[1:], " ")
	reg = regexp.MustCompile(`(?P<first>\s)*\-{1,2}p(?:ort)?(?:\s|=)?(\:\d+)(?P<last>\s)*`)
	submatch := reg.FindStringSubmatch(args)
	if len(submatch) == 0 {
		fmt.Println("You must enter port correctly!")
		fmt.Println("Like -p :80, --p=:80, --port=:8080 or --port :8080")
		return
	}

	port = submatch[2]
	first := submatch[1] == " "
	last := submatch[3] == " "

	switch {
	case first && last:
		rawMsg = strings.Join(reg.Split(args, -1), " ")
	case first, last:
		rawMsg = strings.Join(reg.Split(args, -1), "")
	}
	if len([]byte(rawMsg)) == 0 {
		fmt.Println("You must enter the message you want to send!")
		return
	}
	msg = []byte(rawMsg)
}

func check(err error) {
	if err != nil {
		panic("Error in client:" + err.Error())
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write(msg)
	check(err)

	// receive message echoed from server
	// but why conn.Read gets stuck in a for loop
	// even check io.EOF
	fmt.Print("Echo: ")
	buf := make([]byte, len(msg))
	n, err := conn.Read(buf)
	check(err)
	fmt.Printf("%s\n", buf[:n])

}

func main() {
	// sem := make(chan int)
	conn, err := net.Dial("tcp", "localhost"+port)
	check(err)

	handleConn(conn)
}
