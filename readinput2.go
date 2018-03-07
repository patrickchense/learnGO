package main

import (
	"bufio"
	"os"
	"fmt"
)

var inputReader *bufio.Reader
var input2 string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input2, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input2)
	}
}
/*
Please enter some input:
Chen zhe input
The input was: Chen zhe input
 */