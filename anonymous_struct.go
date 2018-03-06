package main

import "fmt"

type Astruct struct {
	f float64
	int
	string
}

func main()  {
	a := new(Astruct)
	a.f = 1.2
	a.int = 1
	a.string ="test"
	fmt.Printf("a : %v", a)
}

// a : &{1.2 1 test}