package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	inputFile := "input.dat"
	outputFile := "input_copy.dat"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
