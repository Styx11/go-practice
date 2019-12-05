package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// unmarshal json into a known struct
func main() {
	file, err := os.Open("json_shared.json")
	check(err)
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	check(err)

	var vc vCard
	err = json.Unmarshal(buf, &vc)
	check(err)

	vAddress := vc.Address
	fmt.Printf("FristName: %s\n", vc.FirstName)
	fmt.Printf("LastName: %s\n", vc.LastName)
	fmt.Print("Address: \n")
	for k, addr := range vAddress {
		fmt.Println("")
		fmt.Printf("  Type: %s\n", addr.Type)
		fmt.Printf("  City: %s\n", addr.City)
		fmt.Printf("  Country: %s\n", addr.Country)
		if k == len(vAddress)-1 {
			fmt.Println("")
		}
	}
	fmt.Printf("Remark: %v\n", vc.Remark)
}
