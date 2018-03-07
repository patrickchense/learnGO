package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
/*
The input was: asdfsdafasdfsd
 */