package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type address struct {
	Type    string
	City    string
	Country string
}

type vCard struct {
	FirstName string
	LastName  string
	Address   []*address
	Remark    bool
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// defined a vCark struct and marshal to json encoding
func main() {
	pa := &address{"private", "Beijin", "China"}
	wa := &address{"word", "Shanghai", "China"}
	vc := vCard{"Burce", "Lee", []*address{pa, wa}, false}

	js, err := json.MarshalIndent(vc, "", "  ")
	check(err)

	fmt.Fprintf(os.Stdout, "%s\n", js)
}
