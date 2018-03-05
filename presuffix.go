package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "This is an example of a string"
	fmt.Printf("T/F ? Does the string \"%s\" have prefix %s ?", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}

// T/F ? Does the string "This is an example of a string" have prefix Th ?true

