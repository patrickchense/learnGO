package main

import "fmt"

func main()  {
	s := []byte{'p', 'o', 'e', 'm'}
	s1 := s[2:]
	fmt.Println(s1)
	s1[1] = 't'
	fmt.Println(s1)
	fmt.Println(s)
}
/*
[101 109]
[101 116]
[112 111 101 116]
 */