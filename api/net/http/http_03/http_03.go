package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var scheme *regexp.Regexp

func init() {
	scheme = regexp.MustCompile(`^[a-z]+\:\/\/`)
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func receiveURL() []string {
	var urls []string
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter urls you want head to (we won't check their effectiveness)")
	fmt.Println("We may add scheme part of the url you enter, default 'http://'")
	fmt.Println("-------------*-------------")

	for {
		fmt.Print("Enter the url (press Q to quit): ")
		input, err := inputReader.ReadString('\n')
		check(err)

		input = strings.Trim(input, "\n")
		if input == "Q" {
			fmt.Println()
			return urls
		}

		if !scheme.MatchString(input) {
			input = "http://" + input
		}

		urls = append(urls, input)
	}
}

func urlHeader(urls []string) <-chan string {
	response := make(chan string)

	for i, url := range urls {
		go func(i int, url string) {
			res, err := http.Head(url)
			check(err)

			msg := fmt.Sprintf("Head status of %s: %s", url, res.Status)
			response <- msg
		}(i, url)
	}

	return response
}

func main() {
	urls := receiveURL()
	if len(urls) == 0 {
		return
	}

	response := urlHeader(urls)

	for i := 0; i < len(urls); i++ {
		res := <-response
		fmt.Println(res)
	}
}
