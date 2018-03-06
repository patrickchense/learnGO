package main

import "fmt"

var s []int
func main() {
	s = []int{1, 2, 3}
	fmt.Println("The length of s before enlarging is:", len(s))
	fmt.Println(s)
	s = enlarge(s, 5)
	fmt.Println("The length of s after enlarging is:", len(s))
	fmt.Println(s)
}
func enlarge(s []int, factor int) []int{
	l := len(s)
	nl := factor * l
	ns := make([]int, nl)
	copy(ns, s)
	s = ns
	return s
/*	ns := make([]int, len(s)*factor)
	// fmt.Println("The length of ns  is:", len(ns))
	copy(ns, s)
	//fmt.Println(ns)
	s = ns
	//fmt.Println(s)
	//fmt.Println("The length of s after enlarging is:", len(s))
	return s*/
}
/*
The length of s before enlarging is: 3
[1 2 3]
The length of s after enlarging is: 15
[1 2 3 0 0 0 0 0 0 0 0 0 0 0 0]
 */