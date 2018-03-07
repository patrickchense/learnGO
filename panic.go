package main

import "fmt"

func main() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}
/*
Starting the program
panic: A severe error occurred: stopping the program!

goroutine 1 [running]:
main.main()
	/usr/local/go/src/learnGO/panic.go:7 +0x83
 */