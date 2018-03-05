package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "The quick brown fox jumps over the lazy dog"

	s1 := strings.Fields(str) //按空格

	fmt.Printf("Splitted in slice: %v\n", s1)

	for _, val := range s1 {
		fmt.Printf("%s -", val)
	}

	fmt.Println()

	str2 := "GO1|The ABC of Go|25"

	s2 := strings.Split(str2, "|")

	fmt.Printf("Splitted in slice: %v\n", s2)

	for _, val := range s2 {
		fmt.Printf("%s - ", val)
	}

	fmt.Println()

	s3 := strings.Join(s2, ";")

	fmt.Printf("sl2 joined by ;: %s\n", s3)

}

/*
Splitted in slice: [The quick brown fox jumps over the lazy dog]
The -quick -brown -fox -jumps -over -the -lazy -dog -
Splitted in slice: [GO1 The ABC of Go 25]
GO1 - The ABC of Go - 25 -
sl2 joined by ;: GO1;The ABC of Go;25
 */