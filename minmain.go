package main

import (
	"learnGO/min"
	"fmt"
)

func ints1() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := min.IntArray(data) //conversion to type IntArray
	m := min.Min(a)
	fmt.Printf("The minimum of the array is: %v\n", m)
}

func strings1() {
	data := []string{"ddd", "eee", "bbb", "ccc", "aaa"}
	a := min.StringArray(data)
	m := min.Min(a)
	fmt.Printf("The minimum of the array is: %v\n", m)
}

func main() {
	ints1()
	strings1()
}
/*
The minimum of the array is: -5467984
The minimum of the array is: aaa
 */