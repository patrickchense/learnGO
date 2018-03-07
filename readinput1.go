package main

import "fmt"

var (
	firstName, lastName, s1 string
	i1 int
	f3 float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)

func main()  {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f3, &i1, &s1)
	fmt.Println("From the string we read: ", f3, i1, s1)

}
/*
Please enter your full name:
Chen zhe
Hi Chen zhe!
From the string we read:  56.12 5212 Go
 */