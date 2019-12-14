package main

import (
	"encoding/json"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	inputFile, err := os.Open("./json_shared.json")
	check(err)
	defer inputFile.Close()

	outputFile, err := os.OpenFile("./json_output.json", os.O_WRONLY|os.O_CREATE, 0666)
	check(err)
	defer outputFile.Close()

	var vc interface{}
	decoder := json.NewDecoder(inputFile)
	decoder.Decode(&vc)

	encoder := json.NewEncoder(outputFile)
	encoder.Encode(vc)

	os.Stdout.Write([]byte("Done!\n"))
}
