package main

import (
	"fmt"
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
