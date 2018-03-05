package main

import "fmt"

func main()  {
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s) // prints same string
}

/*
Here is the pointer p: 0xc42000e1d0
Here is the string *p: ciao
Here is the string s: ciao
 */