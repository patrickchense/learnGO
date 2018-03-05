package main

import (
	"fmt"
	"strings"
)

func main()  {
	var origS string = "Hi there!"
	var newS string
	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeat string is : %s \n", newS)
}
// The new repeat string is : Hi there!Hi there!Hi there!
