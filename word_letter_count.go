package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
var cc,cw,cl int = 0, 0, 0
func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop: ")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		if input == "S\n" { // Windows, on Linux it is "S\n"
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", cc)
			fmt.Printf("Number of words: %d\n", cw)
			fmt.Printf("Number of lines: %d\n", cl)
			os.Exit(0)
		}
		Counters(input)
	}
}

func Counters(input string) {
	cc += len(input) - 1 // -2 for \n
	cw += len(strings.Fields(input))
	cl++
}
