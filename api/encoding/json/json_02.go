package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// unmarshal json object and write it to console
func main() {
	var v interface{}
	f, err := os.Open("./json_shared.json")
	check(err)
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	check(err)

	err = json.Unmarshal(buf, &v)
	check(err)

	vCard := v.(map[string]interface{})
	for k, vv := range vCard {
		switch vc := vv.(type) {
		case string:
			fmt.Printf("%s: %s\n", k, vc)
		case []interface{}:
			fmt.Printf("%s:\n", k)

			for _, v := range vc {
				ut := v.(map[string]interface{})
				for uk, addr := range ut {
					fmt.Print("  ")
					fmt.Printf("%s: %s\n", uk, addr.(string))
				}
				fmt.Println("")
			}

		}
	}
}
