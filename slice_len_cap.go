package main

import "fmt"

func main()  {
	s := make([]int, 5)
	fmt.Printf("s len:%d, cap:%d", len(s), cap(s))
	s = s[2:4]
	fmt.Println()
	fmt.Printf("s len:%d, cap:%d", len(s), cap(s))

}
/*
s len:5, cap:5
s len:2, cap:3  cap = 5 - 2 = 3  2 is index 0, 1 in array
 */